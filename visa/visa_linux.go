package visa

/*
#cgo linux CFLAGS: -I/usr/include/rsvisa
#cgo linux LDFLAGS: -lrsvisa
#include <stdlib.h> //for free()
#include <stdint.h>
#include <visa.h>

inline static uint32_t write(ViSession s, _GoString_ cmd){
	ViUInt32 ret;
	viWrite(s, (ViBuf)_GoStringPtr(cmd), _GoStringLen(cmd), &ret);
	return (uint32_t)ret;
}

inline static uint32_t query(ViSession s, _GoString_ cmd, uint8_t* resp, int len){
	ViUInt32 ret = write(s, cmd);
	if (ret == 0){
		return len;
	}
	viRead(s, resp, len, &ret);
	return (uint32_t)ret;
}
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
	// convert to C string, which is ends with '\0'
	name := C.CString(res)
	r := C.viOpen(rm, (*C.ViChar)(unsafe.Pointer(name)), C.VI_LOAD_CONFIG, C.ViUInt32(timeout), &s)

	// need free the string
	C.free(unsafe.Pointer(name))

	if r != C.VI_SUCCESS{
		return s, errors.New(fmt.Sprintf("viOpen error, code is %d", r))
	}
	return s, err
}

func (s C.ViSession) Close(){
	C.viClose(C.ViObject(s))
}

func (s C.ViSession) Write(cmd string) uint32{
	return uint32(C.write(s, cmd))
}

func (s C.ViSession) Query(cmd string) string{
	resp := make([]byte, 128)

	C.query(s, cmd, (*C.uint8_t)(&resp[0]), C.int(cap(resp)))
	return string(resp)
}