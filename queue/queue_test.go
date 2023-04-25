package queue_test

import (
	"testing"

	"github.com/nellystanford/golang/queue"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Should be able to create a new queue", func(t *testing.T) {
		// Prepare

		// Act
		q := queue.NewQueue()

		// Assert
		assert.NotNil(t, q)
	})
	t.Run("Newly created queue should be empty", func(t *testing.T) {
		// Prepare

		// Act
		q := queue.NewQueue()

		// Assert
		assert.Equal(t, 0, q.Len())
	})

	t.Run("Should be able to insert items to the queue and change its length", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")

		// Assert
		assert.Equal(t, 1, q.Len())
	})

	t.Run("Should be able to insert 2 items to the queue and change its length to two", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")
		q.Insert("The Devil Wears Prada")

		// Assert
		assert.Equal(t, 2, q.Len())
	})

	t.Run("Should be able to insert 2 items to the queue, remove the one on the top and change the stack length", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")
		q.Insert("The Devil Wears Prada")
		q.Remove()

		// Assert
		assert.Equal(t, 1, q.Len())
	})

	// t.Run("Should be able to push 2 items to the stack and pop one the last of them", func(t *testing.T) {
	// 	// Prepare
	// 	s := stack.NewStack()

	// 	// Act
	// 	s.Push(100)
	// 	s.Push(50)
	// 	result := s.Pop()

	// 	// Assert
	// 	assert.Equal(t, 1, s.Len())
	// 	assert.Equal(t, 50, result)
	// })

	// t.Run("Should be able to push 2 items and then pop 2 items", func(t *testing.T) {
	// 	// Prepare
	// 	s := stack.NewStack()

	// 	// Act
	// 	s.Push(100)
	// 	s.Push(50)
	// 	result1 := s.Pop()
	// 	result2 := s.Pop()

	// 	// Assert
	// 	assert.Equal(t, 0, s.Len())
	// 	assert.Equal(t, 50, result1)
	// 	assert.Equal(t, 100, result2)
	// })
}
