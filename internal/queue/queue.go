package queue

type Queue struct {
	SnakeBody []any
}

func CreateQueue() Queue {
	return Queue {
		SnakeBody: make([]any, 1),
	}
}

func (q *Queue) Append(currentCords []string) {
	q.SnakeBody = append(q.SnakeBody, currentCords)
}

func (q *Queue) Pop() {
	q.SnakeBody = q.SnakeBody[:len(q.SnakeBody) - 1]
}

func (q *Queue) Peek() []any {
	return []any{q.SnakeBody[0]}
}

func (q *Queue) AddToHead() {
	head := q.Peek()
	q.SnakeBody = append(head, q.SnakeBody)
}
















