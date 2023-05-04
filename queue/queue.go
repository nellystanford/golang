package queue

// implementar sem tail e sem head, usar indices do slice

type Queue struct {
	//tail   int // zero value is 0
	//head   int
	movies []string
}

func NewQueue() *Queue {
	//return &Queue{tail: -1, head: -1, movies: make([]string, 10)}
	return &Queue{movies: []string{}}
}

func (q *Queue) Len() int {
	return len(q.movies)
}

func (q *Queue) Insert(value string) {
	q.movies = append(q.movies, value)
}

func (q *Queue) Remove() int {
	q.movies = q.movies[1:]
	return q.Len()
}

func (q *Queue) NextMovie() string {
	return q.movies[0]
}
