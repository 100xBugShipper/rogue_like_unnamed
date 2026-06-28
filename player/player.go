package player

import (
	"log"
	"time"

	"github.100xBugShipper/rogue_like/queue"
)

type Snake struct {
	Symbol     string
	X          int
	Y          int
	FoodCount  int
	SnakeQueue queue.Queue
}

func CreateSnake() *Snake {
	randRow := 10
	randCol := 10

	return &Snake{
		Symbol:     "@",
		X:          randRow,
		Y:          randCol,
		FoodCount:  0,
		SnakeQueue: queue.CreateQueue(),
	}
}

func (snk *Snake) AutoMove(moveChan chan string, lastMove string) string {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		_, ok := <-ticker.C
		if ok {
			log.Println("I am sending the last move into the channel")
			moveChan <- lastMove
		}
	}
}

func (snk *Snake) Eat() {
}

func (snk *Snake) Grow() {
}




