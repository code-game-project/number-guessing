package numberguessing

import (
	"math/rand"

	"github.com/code-game-project/go-server/cg"
)

type Game struct {
	cg     *cg.Game
	number int
	tries  int
	min    int
	max    int
}

func NewGame(cgGame *cg.Game, config GameConfig) *Game {
	game := &Game{
		cg:     cgGame,
		number: rand.Intn(config.Max+1-config.Min) + config.Min,
		min:    config.Min,
		max:    config.Max,
	}
	return game
}

func (g *Game) Run() {
	for g.cg.Running() {
		cmd, ok := g.cg.WaitForNextCommand()
		if !ok {
			break
		}
		g.handleEvent(cmd.Origin, cmd.Cmd)
	}
}

func (g *Game) handleEvent(origin *cg.Player, cmd cg.Command) {
	switch cmd.Name {
	case GuessCmd:
		g.tries++
		var data GuessCmdData
		err := cmd.UnmarshalData(&data)
		if err != nil {
			origin.Log.Error("Failed to unmarshal command data: %s", err)
			return
		}
		if data.Number > g.number {
			g.cg.Log.InfoData(cmd, "The user guessed too high. (%d tries)", g.tries)
			g.cg.Send(TooHighEvent, nil)
		} else if data.Number < g.number {
			g.cg.Log.InfoData(cmd, "The user guessed too low. (%d tries)", g.tries)
			g.cg.Send(TooLowEvent, nil)
		} else {
			g.cg.Send(CorrectEvent, CorrectEventData{
				Number: g.number,
				Tries:  g.tries,
			})
			g.tries = 0
			g.number = rand.Intn(g.max+1-g.min) + g.min
			g.cg.Log.InfoData(cmd, "The user guessed the correct number (%d) after %d tries.", g.number, g.tries)
		}
	}
}
