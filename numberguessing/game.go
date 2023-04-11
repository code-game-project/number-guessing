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
		type guessLogMessage struct {
			Number int `json:"number"`
			Tries  int `json:"tries"`
		}
		if data.Number > g.number {
			g.cg.Log.InfoData(guessLogMessage{
				Number: data.Number,
				Tries:  g.tries,
			}, "The user guessed too high.")
			g.cg.Send(TooHighEvent, TooHighEventData(data))
		} else if data.Number < g.number {
			g.cg.Log.InfoData(guessLogMessage{
				Number: data.Number,
				Tries:  g.tries,
			}, "The user guessed too low.")
			g.cg.Send(TooLowEvent, TooLowEventData(data))
		} else {
			g.cg.Send(CorrectEvent, CorrectEventData{
				Number: g.number,
				Tries:  g.tries,
			})
			g.cg.Log.InfoData(cmd, "The user guessed the correct number (%d) after %d tries.", g.number, g.tries)
			g.tries = 0
			g.number = rand.Intn(g.max+1-g.min) + g.min
		}
	}
}
