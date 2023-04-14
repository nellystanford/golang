package stack

type Stack struct {
	length int // zero value is 0
	items  []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Push(value int) {
	s.length += 1
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	s.length -= 1
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}
