package queue_test

import (
	"testing"

	"github.com/nellystanford/golang/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
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

	t.Run("Should be able to insert 2 items to the queue, remove the one in the beginning and change the queue length", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")
		q.Insert("The Devil Wears Prada")
		q.Remove()

		// Assert
		assert.Equal(t, 1, q.Len())
	})

	t.Run("Should be able to insert 2 items to the queue, remove the first two and change the queue length", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")
		q.Insert("The Devil Wears Prada")
		q.Insert("The Age of Adeline")
		q.Remove()
		q.Remove()

		// Assert
		assert.Equal(t, 1, q.Len())
	})

	t.Run("Should be able to insert 2 items to the queue, remove the first two, change the queue length and confirm the item left", func(t *testing.T) {
		// Prepare
		q := queue.NewQueue()

		// Act
		q.Insert("About Time")
		q.Insert("The Devil Wears Prada")
		q.Insert("The Age of Adaline")
		len1 := q.Remove()
		len2 := q.Remove()

		// Assert
		assert.Equal(t, "The Age of Adaline", q.NextMovie())
		assert.Equal(t, 2, len1)
		assert.Equal(t, 1, len2)
	})
}
