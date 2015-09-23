package cp

// #cgo CFLAGS: -DNDEBUG
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
	C.cpSpaceSetGravity(s.ptr, C.cpVect{C.CGFloat(x), C.CGFloat(y)})
}
