package ipp

/*
#include <ipp.h>
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

func Version() string {
	v := C.ippsGetLibVersion()
	return fmt.Sprintf("%d.%d.%d", int(v.major), int(v.minor), int(v.majorBuild))
}

func checkStatus(s C.IppStatus) {
	if s != C.ippStsNoErr {
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()

		panic(fmt.Sprintf("MKL error, File:%s, Line:%d, Func:%s, %s",
			frame.File, frame.Line, frame.Function, string(C.GoString(C.ippGetStatusString(s)))))
	}
}

// Copies the contents of one vector into another
func Copy(src, dst interface{}) {
}

// Initializes a vector to zero
func Zero(dst interface{}) {

	switch dst.(type) {
	case []float32:
		p, _ := dst.([]float32)
		checkStatus(C.ippsZero_32f((*C.Ipp32f)(unsafe.Pointer(&p[0])), C.int(len(p))))

	case []float64:
		p, _ := dst.([]float64)
		checkStatus(C.ippsZero_64f((*C.Ipp64f)(unsafe.Pointer(&p[0])), C.int(len(p))))

	default:
		panic("Not supported type")
	}
}

// Adds a constant value to each element of a vector
// In-place operations
func AddC_I(dst interface{}, val float64) {
	switch dst.(type) {
	case []float32:
		p, _ := dst.([]float32)
		checkStatus(C.ippsAddC_32f_I(C.Ipp32f(val),
			(*C.Ipp32f)(unsafe.Pointer(&p[0])),
			C.int(len(p))))

	case []float64:
		p, _ := dst.([]float64)
		checkStatus(C.ippsAddC_64f_I(C.Ipp64f(val),
			(*C.Ipp64f)(unsafe.Pointer(&p[0])),
			C.int(len(p))))

	default:
		panic("Not supported type")
	}
}

// Adds a constant value to each element of a vector
// Not-in-place operations
func AddC(dst interface{}, val float64) interface{} {
	switch dst.(type) {
	case []float32:
		p, _ := dst.([]float32)
		r := make([]float32, cap(p), len(p))
		checkStatus(C.ippsAddC_32f((*C.Ipp32f)(unsafe.Pointer(&p[0])),
			C.Ipp32f(val),
			(*C.Ipp32f)(unsafe.Pointer(&r[0])),
			C.int(len(p))))
		return r

	case []float64:
		p, _ := dst.([]float64)
		r := make([]float64, cap(p), len(p))
		checkStatus(C.ippsAddC_64f((*C.Ipp64f)(unsafe.Pointer(&p[0])),
			C.Ipp64f(val),
			(*C.Ipp64f)(unsafe.Pointer(&r[0])),
			C.int(len(p))))
		return r

	default:
		panic("Not supported type")
	}
	return nil
}
