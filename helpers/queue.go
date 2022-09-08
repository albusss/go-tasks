package helpers

type Queue struct {
	Q []interface{}
}

func (q *Queue) Enqueue(val interface{}) {
	q.Q = append(q.Q, val)
}

func (q *Queue) Denqueue() interface{} {
	val := q.Q[0]
	q.Q = q.Q[1:]

	return val
}
