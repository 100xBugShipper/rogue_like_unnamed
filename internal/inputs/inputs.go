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

func inputChanged(checkInp chan string) {
	var inputChanged = ""
	fmt.Scan(&inputChanged)

	checkInp <- inputChanged
}

func (pi *PlayerInput) DetectKeys() {
	var playerInput string
	checkInp := make(chan string, 1)
	fmt.Println("<-- |>Game Started<| -->")

	for playerInput != "q" {
		fmt.Scan(&playerInput)

		for {
			pi.MoveChan <- playerInput
			go inputChanged(checkInp)
			a := <- checkInp
			if a != playerInput {
				break
			}
		}
	}

	close(pi.MoveChan)
}
