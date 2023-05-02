package lib_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("Should be able to air local server", func(t *testing.T) {
		// Prepare

		// Act
		n := lib.LocalHost()

		// Assert
		assert.NotNil(t, n)
	})
}
