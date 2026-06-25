package playerInputs

import (
	"fmt"
)

type PlayerInput struct {
	MoveChan chan string
}

func CreatePlayerInputObj() *PlayerInput {
	return &PlayerInput{
		MoveChan: make(chan string),
	}
}

func (pi *PlayerInput) DetectKeys() {
	var playerInput string
	fmt.Println("<-- |>Game Started<| -->")

	for playerInput != "q" {
		fmt.Scan(&playerInput)

		pi.MoveChan <- playerInput
	}
}
