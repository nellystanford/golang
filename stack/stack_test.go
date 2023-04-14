package stack_test

import (
	"testing"

	"github.com/nellystanford/golang/stack"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Should be able to create a new stack", func(t *testing.T) {
		// Prepare

		// Act
		s := stack.NewStack()

		// Assert
		assert.NotNil(t, s)
	})
	t.Run("Newly created stack should be empty", func(t *testing.T) {
		// Prepare

		// Act
		s := stack.NewStack()

		// Assert
		assert.Equal(t, 0, s.Len())
	})

	t.Run("Should be able to push items to the stack and change its length", func(t *testing.T) {
		// Prepare
		s := stack.NewStack()

		// Act
		s.Push(100)

		// Assert
		assert.Equal(t, 1, s.Len())
	})
	t.Run("Should be able to push 2 items to the stack and change its length to two", func(t *testing.T) {
		// Prepare
		s := stack.NewStack()

		// Act
		s.Push(100)
		s.Push(50)

		// Assert
		assert.Equal(t, 2, s.Len())
	})

	t.Run("Should be able to push 2 items to the stack, pop one of them and change the stack length", func(t *testing.T) {
		// Prepare
		s := stack.NewStack()

		// Act
		s.Push(100)
		s.Push(50)
		s.Pop()

		// Assert
		assert.Equal(t, 1, s.Len())
	})

	t.Run("Should be able to push 2 items to the stack and pop one the last of them", func(t *testing.T) {
		// Prepare
		s := stack.NewStack()

		// Act
		s.Push(100)
		s.Push(50)
		result := s.Pop()

		// Assert
		assert.Equal(t, 1, s.Len())
		assert.Equal(t, 50, result)
	})

	t.Run("Should be able to push 2 items and then pop 2 items", func(t *testing.T) {
		// Prepare
		s := stack.NewStack()

		// Act
		s.Push(100)
		s.Push(50)
		result1 := s.Pop()
		result2 := s.Pop()

		// Assert
		assert.Equal(t, 0, s.Len())
		assert.Equal(t, 50, result1)
		assert.Equal(t, 100, result2)
	})
}
