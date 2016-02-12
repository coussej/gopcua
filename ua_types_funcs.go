/*
  Copyright 2016 Jeroen Coussement.
*/

package gopcua

/*
#cgo CFLAGS: -std=c99 -I .
#include <stdlib.h>
#include "gopcua_client.h"
*/
import "C"
import "unsafe"

/****************************/
/* Builtin Type Definitions */
/****************************/

/********************************************/
/* String: A sequence of Unicode characters */
/********************************************/

func UA_STRING(s string) (uas UA_String) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	r := C.UA_STRING(n)
	return *(*UA_String)(unsafe.Pointer(&r))
}

/*********************************/
/* DateTime: An instance in time */
/*********************************/

/* The current time */
func UA_DateTime_now() UA_DateTime {
	n := C.UA_DateTime_now()
	return *(*UA_DateTime)(unsafe.Pointer(&n))
}

/* CPU clock invariant to system time changes. Use only for time diffs, not current time */
func UA_DateTime_nowMonotonic() UA_DateTime {
	n := C.UA_DateTime_nowMonotonic()
	return *(*UA_DateTime)(unsafe.Pointer(&n))
}

func UA_DateTime_toStruct(dt UA_DateTime) UA_DateTimeStruct {
	cdt := *(*C.UA_DateTime)(&dt)
	dts := C.UA_DateTime_toStruct(cdt)
	return *(*UA_DateTimeStruct)(unsafe.Pointer(&dts))
}

func UA_DateTime_toString(dt UA_DateTime) UA_String {
	cdt := *(*C.UA_DateTime)(&dt)
	dts := C.UA_DateTime_toString(cdt)
	return *(*UA_String)(unsafe.Pointer(&dts))
}

/**************************************************************************/
/* Guid: A 16 byte value that can be used as a globally unique identifier */
/**************************************************************************/

func UA_Guid_random() UA_Guid {
	g := C.UA_Guid_random()
	return *(*UA_Guid)(unsafe.Pointer(&g))
}
