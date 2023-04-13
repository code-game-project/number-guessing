package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/spf13/pflag"

	"github.com/code-game-project/number-guessing/numberguessing"
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	var port int
	pflag.IntVarP(&port, "port", "p", 0, "The network port of the game server.")
	var frontend string
	pflag.StringVar(&frontend, "frontend", "", "The root directory of the frontend")
	pflag.Parse()

	if port == 0 {
		portStr, ok := os.LookupEnv("CG_PORT")
		if ok {
			port, _ = strconv.Atoi(portStr)
		}
	}

	if port == 0 {
		port = 8080
	}

	server := cg.NewServer("number-guessing", cg.ServerConfig{
		DisplayName:             "Number Guessing",
		MaxPlayersPerGame:       1,
		DeleteInactiveGameDelay: 30 * time.Second,
		KickInactivePlayerDelay: 5 * time.Minute,
		Version:                 "0.2",
		Description:             "A number guessing game.",
		RepositoryURL:           "https://github.com/code-game-project/number-guessing",
		Port:                    port,
		CGEFilepath:             "events.cge",
		WebRoot:                 frontend,
	})

	server.Run(func(cgGame *cg.Game, config json.RawMessage) {
		var gameConfig numberguessing.GameConfig
		err := json.Unmarshal(config, &gameConfig)
		if err != nil || gameConfig.Min > gameConfig.Max {
			cgGame.Log.ErrorData(config, "invalid game config")
			gameConfig.Min = 0
			gameConfig.Max = 100
		}
		if gameConfig.Max == 0 {
			gameConfig.Max = 100
		}

		cgGame.SetConfig(gameConfig)

		numberguessing.NewGame(cgGame, gameConfig).Run()
	})
}
