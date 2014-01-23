package crypt

import (
	"unsafe"
)

// #cgo LDFLAGS: -lcrypt
// #define _GNU_SOURCE
// #include <crypt.h>
// #include <stdlib.h>
import "C"

// Crypt wraps C library crypt_r function
func Crypt(key, salt string) string {
	data := C.struct_crypt_data{}
	ckey := C.CString(key)
	csalt := C.CString(salt)
	out := C.GoString(C.crypt_r(ckey, csalt, &data))
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(csalt))
	return out
}
