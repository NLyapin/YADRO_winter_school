package queue

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
