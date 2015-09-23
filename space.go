package cp

// #cgo CFLAGS: -DNDEBUG -std=gnu99
// #cgo linux LDFLAGS: -lm
//
// #define CP_USE_CGTYPES 0
// #include "chipmunk/chipmunk.h"
import "C"

type Space struct {
	ptr *C.cpSpace
}

// Allocate and initialize a Space.
func NewSpace() *Space {
	space := new(Space)
	space.ptr = C.cpSpaceNew()
	return space
}

// Get the gravity passed to rigid bodies when integrating velocity.
func (s *Space) GetGravity() (float64, float64) {
	vect := C.cpSpaceGetGravity(s.ptr)
	return float64(vect.x), float64(vect.y)
}

// Set the gravity to pass to rigid bodies when integrating velocity.
func (s *Space) SetGravity(x float64, y float64) {
	C.cpSpaceSetGravity(s.ptr, C.cpVect{C.cpFloat(x), C.cpFloat(y)})
}
