package cp

import (
	// "fmt"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestGravity(t *testing.T) {
	space := NewSpace()
	gravityX, gravityY := -10.0, 0.0

	space.SetGravity(gravityX, gravityY)
	x, y := space.GetGravity()

	assert.Equal(t, gravityX, x)
	assert.Equal(t, gravityY, y)
}
