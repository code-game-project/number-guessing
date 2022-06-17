package numberguessing

import (
	"fmt"

	"github.com/code-game-project/go-server/cg"
)

type Game struct {
	cg *cg.Game
}

func NewGame(cgGame *cg.Game) *Game {
	game := &Game{
		cg: cgGame,
	}
	return game
}

func (g *Game) Run() {
	for g.cg.Running() {
		event, ok := g.cg.WaitForNextEvent()
		if !ok {
			break
		}
		g.handleEvent(event.Player, event.Event)
	}
}

func (g *Game) handleEvent(player *cg.Player, event cg.Event) {
	switch event.Name {
	default:
		player.Send(player.Id, cg.ErrorEvent, cg.ErrorEventData{
			Message: fmt.Sprintf("unexpected event: %s", event.Name),
		})
	}
}