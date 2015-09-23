package cp

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestEquals(t *testing.T) {
	vec1 := Vect{10, 20}
	vec2 := Vect{10, 20}
	vec3 := Vect{30, 40}

	assert.True(t, vec1.Equals(vec2))
	assert.False(t, vec1.Equals(vec3))
}

func TestAdd(t *testing.T) {
	vec1 := Vect{10, 20}
	vec2 := Vect{30, 50}
	vec3 := Vect{40, 70}

	assert.Equal(t, vec1.Add(vec2), vec3)
}

func TestSub(t *testing.T) {
	vec1 := Vect{30, 80}
	vec2 := Vect{10, 50}
	vec3 := Vect{20, 30}

	assert.Equal(t, vec1.Sub(vec2), vec3)
}

func TestNeg(t *testing.T) {
	vec1 := Vect{10, 20}
	vec2 := Vect{-10, -20}

	assert.Equal(t, vec1.Neg(), vec2)
}

func TestMult(t *testing.T) {
	vec1 := Vect{30, 20}
	vec2 := Vect{60, 40}

	assert.Equal(t, vec1.Mult(2), vec2)
}

func TestDot(t *testing.T) {
	vec1 := Vect{30, 20}
	vec2 := Vect{10, 50}
	dot := 1300.0

	assert.Equal(t, vec1.Dot(vec2), dot)
}
