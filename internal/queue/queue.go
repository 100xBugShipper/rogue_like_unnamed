package queue

type Queue struct {
	SnakeBody []Cords
}

type Cords struct {
	X int
	Y int
}

func CreateCords() *Cords {
	return &Cords{
		X: 0,
		Y: 0,
	}
}

func CreateQueue() Queue {
	return Queue {
		SnakeBody: make([]Cords, 1),
	}
}

func (q *Queue) Append(cords Cords) {
	q.SnakeBody = append(q.SnakeBody, cords)
}














