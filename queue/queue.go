package queue

type Queue struct {
	size   int // zero value is 0
	movies []string
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Insert(value string) {
	q.size += 1
	q.movies = append(q.movies, value)
}

func (q *Queue) Remove() string {
	q.size -= 1
	result := q.movies[0]
	q.movies = q.movies[1:]
	return result
}
