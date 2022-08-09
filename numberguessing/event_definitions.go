/*
A number guessing game.
*/
package numberguessing

import "github.com/code-game-project/go-server/cg"

type GameConfig struct {
	// The minimum number (inclusive) default=0
	Min int `json:"min"`
	// The maximum number (inclusive) default=100
	Max int `json:"max"`
}

// Send the `guess` command to make a guess.
const GuessCmd cg.CommandName = "guess"

type GuessCmdData struct {
	// The number you guess.
	Number int `json:"number"`
}

// You guessed too high.
const TooHighEvent cg.EventName = "too_high"

type TooHighEventData struct{}

// Your too low.
const TooLowEvent cg.EventName = "too_low"

type TooLowEventData struct{}

// You guessed the correct number.
const CorrectEvent cg.EventName = "correct"

type CorrectEventData struct {
	// The correct number.
	Number int `json:"number"`
	// The number of tries.
	Tries int `json:"tries"`
}
