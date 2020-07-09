package visa

/*
#cgo linux CFLAGS: -I/usr/include/rsvisa
#cgo linux LDFLAGS: -lrsvisa
#include <visa.h>
 */
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)


/* global resource manager */
var rm C.ViSession

func init(){
	C.viOpenDefaultRM(&rm)
}

func Open(res string, timeout int32) (s C.ViSession, err error){
	r := C.viOpen(rm, (*C.ViChar)(unsafe.Pointer(C.CString(res))), 4, C.ViUInt32(timeout), &s)
	if r != C.VI_SUCCESS{
		return s, errors.New(fmt.Sprintf("error is %d", r))
	}
	return s, err
}

func (s C.ViSession) Close(){
	C.viClose(C.ViObject(s))
}

func (s C.ViSession) Write(cmd string) int32{
	var retCnt int32
	C.viWrite(s, (*C.ViByte)(unsafe.Pointer(C.CString(cmd))), C.ViUInt32(len(cmd)), (*C.ViUInt32)(unsafe.Pointer(&retCnt)))
	return retCnt
}

func (s C.ViSession) Query(cmd string) string{
	s.Write(cmd)
	var retCnt int32
	var buf = make([]byte, 1024)
	C.viRead(s, (*C.ViByte)(unsafe.Pointer(&buf[0])), 1024, (*C.ViUInt32)(unsafe.Pointer(&retCnt)))
	return string(buf)
}