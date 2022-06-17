package main

import (
	"os"
	"strconv"
	"time"

	"github.com/code-game-project/go-server/cg"
	"github.com/code-game-project/number_guessing/numberguessing"
	"github.com/spf13/pflag"
)

func main() {
	var port int
	pflag.IntVarP(&port, "port", "p", 0, "The network port of the game server.")
	pflag.Parse()

	if port == 0 {
		portStr, ok := os.LookupEnv("CG_PORT")
		if ok {
			port, _ = strconv.Atoi(portStr)
		}
	}

	if port == 0 {
		port = 80
	}

	server := cg.NewServer("number_guessing", cg.ServerConfig{
		DisplayName:             "Number Guessing",
		MaxPlayersPerGame:       1,
		DeleteInactiveGameDelay: 1 * time.Minute,
		KickInactivePlayerDelay: 1 * time.Minute,
		Version:                 "0.1",
		Description:             "A number guessing game.",
		RepositoryURL:           "https://github.com/code-game-project/number_guessing",
		Port:                    port,
		CGEFilepath:             "events.cge",
	})

	server.Run(func(cgGame *cg.Game) {
		numberguessing.NewGame(cgGame).Run()
	})
}
