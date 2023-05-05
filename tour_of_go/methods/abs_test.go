package abs_test

import (
	"testing"

	methods "github.com/nellystanford/golang/tour_of_go/methods"
	"github.com/stretchr/testify/assert"
)

func TestMethods(t *testing.T) {
	t.Run("Should be able to create a new vertex", func(t *testing.T) {
		// Prepare

		// Act
		q := methods.NewVertex()

		// Assert
		assert.NotNil(t, q)
	})

	t.Run("Should be able to insert the x and y value of a vertex", func(t *testing.T) {
		// Prepare
		q := methods.NewVertex()

		// Act
		methods.SetVertexValue(q, 3, 4)

		// Assert
		assert.Equal(t, float64(3), q.X)
		assert.Equal(t, float64(4), q.Y)

	})

	t.Run("Should be able to return the abs of a vertex", func(t *testing.T) {
		// Prepare
		q := methods.NewVertex()
		methods.SetVertexValue(q, 3, 4)

		// Act
		aux := q.Abs()

		// Assert
		assert.Equal(t, float64(5), aux)
	})
}
