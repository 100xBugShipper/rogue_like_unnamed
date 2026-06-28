package snake

import (
	"github.100xBugShipper/rogue_like/internal/queue"
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

func (snk *Snake) Move() {
}

func (snk *Snake) Eat() {
}

func (snk *Snake) Grow() {
}




