package queue

type Queue struct {
	tail   int // zero value is 0
	movies []string
	head   int
}

func NewQueue() *Queue {
	//return &Queue{tail: -1, head: -1, movies: make([]string, 10)}
	return &Queue{tail: -1, head: -1, movies: nil}
}

func (q *Queue) Len() int {
	return (q.tail - q.head)
}

func (q *Queue) Insert(value string) {
	q.movies = append(q.movies, value)
	if q.head == -1 {
		q.head = 0
	}
	q.tail += 1
}

func (q *Queue) Remove() int {
	q.head += 1
	return q.Len()
}

func (q *Queue) NextMovie() string {
	return q.movies[q.head]
}
