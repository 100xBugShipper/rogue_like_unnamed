package playerInputs

import (
	"fmt"
	"sync"
)

type PlayerInput struct {
	Wg       *sync.WaitGroup
	MoveChan chan string
	mu       *sync.Mutex
}

func CreatePlayerInputObj() *PlayerInput {
	return &PlayerInput{
		MoveChan: make(chan string),
	}
}

func (pi *PlayerInput) DetectKeys() {
	var playerInput string
	fmt.Println("<-- |>Game Started<| -->")
	defer pi.Wg.Done()

	for playerInput != "q" {
		pi.mu.Lock()
		defer pi.mu.Unlock()
		fmt.Scan(&playerInput)

		pi.MoveChan <- playerInput
	}
}
