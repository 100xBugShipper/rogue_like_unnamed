package snake

import (
	"github.100xBugShipper/rogue_like/internal/queue"
)

type Snake struct {
	Symbol     string
	X          int
	Y          int
	Food       string
	SnakeQueue queue.Queue
	Direction  string
}

func CreateSnake() *Snake {
	randRow := 10
	randCol := 10

	return &Snake{
		Symbol:     "@",
		X:          randRow,
		Y:          randCol,
		Food:		"$",
		SnakeQueue: queue.CreateQueue(),
	}
}

func (snk *Snake) Grow() {
	var cords queue.Cords
	cords.X = snk.X
	cords.Y = snk.Y
	snk.SnakeQueue.Append(cords)
}

func (snk *Snake) Dequeue() {
	if len(snk.SnakeQueue.SnakeBody) > 0 {
		snk.SnakeQueue.SnakeBody = snk.SnakeQueue.SnakeBody[1:]
	}
}










