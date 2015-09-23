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

func NewSpace() *Space {
	space := new(Space)
	space.ptr = C.cpSpaceNew()
	return space
}

func (s *Space) GetGravity() (float64, float64) {
	vect := C.cpSpaceGetGravity(s.ptr)
	return float64(vect.x), float64(vect.y)
}

func (s *Space) SetGravity(x float64, y float64) {
	C.cpSpaceSetGravity(s.ptr, C.cpVect{C.cpFloat(x), C.cpFloat(y)})
}
