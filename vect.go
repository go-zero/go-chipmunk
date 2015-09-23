package cp

// #include "chipmunk/chipmunk.h"
import "C"

type Vect struct {
	X float64
	Y float64
}

// Check if two vectors are equal. (Be careful when comparing floating point numbers!)
func (v Vect) Equals(other Vect) bool {
	return int(C.cpveql(toC(v), toC(other))) == 1
}

// Add two vectors
func (v Vect) Add(other Vect) Vect {
	return fromC(C.cpvadd(toC(v), toC(other)))
}

// Subtract two vectors.
func (v Vect) Sub(other Vect) Vect {
	return fromC(C.cpvsub(toC(v), toC(other)))
}

// Negate a vector.
func (v Vect) Neg() Vect {
	return fromC(C.cpvneg(toC(v)))
}

// Scalar multiplication.
func (v Vect) Mult(s float64) Vect {
	return fromC(C.cpvmult(toC(v), C.cpFloat(s)))
}

/// Vector dot product.
func (v Vect) Dot(other Vect) float64 {
	return float64(C.cpvdot(toC(v), toC(other)))
}

func toC(v Vect) C.cpVect {
	return C.cpVect{C.cpFloat(v.X), C.cpFloat(v.Y)}
}

func fromC(cv C.cpVect) Vect {
	return Vect{float64(cv.x), float64(cv.y)}
}
