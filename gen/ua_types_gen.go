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

const UA_BUILTIN_TYPES_COUNT uint = 25

/** Boolean: A two-state logical value (true or false) */
type UA_Boolean C.UA_Boolean

const (
	UA_TRUE  = true
	UA_FALSE = false
)

/** SByte: An integer value between -128 and 127 */
type UA_SByte C.UA_SByte

const (
	UA_SBYTE_MAX = 127
	UA_SBYTE_MIN = (-128)
)

/** Byte: An integer value between 0 and 256 */
type UA_Byte C.UA_Byte

const (
	UA_BYTE_MAX = 256
	UA_BYTE_MIN = 0
)

/** Int16: An integer value between -32 768 and 32 767 */
type UA_Int16 C.UA_Int16

const (
	UA_INT16_MAX = 32767
	UA_INT16_MIN = -32768
)

/** UInt16: An integer value between 0 and 65 535 */
type UA_UInt16 C.UA_UInt16

const (
	UA_UINT16_MAX = 65535
	UA_UINT16_MIN = 0
)

/** Int32: An integer value between -2 147 483 648 and 2 147 483 647 */
type UA_Int32 C.UA_Int32

const (
	UA_INT32_MAX = 2147483647
	UA_INT32_MIN = -2147483648
)

/** UInt32: An integer value between 0 and 4 294 967 295 */
type UA_UInt32 C.UA_UInt32

const (
	UA_UINT32_MAX = 4294967295
	UA_UINT32_MIN = 0
)

/** Int64: An integer value between -10 223 372 036 854 775 808 and 9 223 372 036 854 775 807 */
type UA_Int64 C.UA_Int64

const (
	UA_INT64_MAX = int64(9223372036854775807)
	UA_INT64_MIN = int64(9223372036854775808)
)

/** UInt64: An integer value between 0 and 18 446 744 073 709 551 615 */
type UA_UInt64 C.UA_UInt64

const (
	UA_UINT64_MAX = int64(18446744073709551615)
	A_UINT64_MIN  = int64(0)
)

/** Float: An IEEE single precision (32 bit) floating point value */
type UA_Float C.UA_Float

/** Double: An IEEE double precision (64 bit) floating point value */
type UA_Double C.UA_Double

/********************************************/
/* String: A sequence of Unicode characters */
/********************************************/

/** String: A sequence of Unicode characters **/
type UA_String C.UA_String

func UA_STRING(s string) (uas UA_String) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	uas = UA_String(C.UA_STRING(n))
	return
}

/*********************************/
/* DateTime: An instance in time */
/*********************************/

/* A DateTime value is encoded as a 64-bit signed integer which represents the
   number of 100 nanosecond intervals since January 1, 1601 (UTC) */
type UA_DateTime C.UA_DateTime

/* Multiply to convert units for time difference computations */
const (
	UA_USEC_TO_DATETIME = int64(14)
	UA_MSEC_TO_DATETIME = UA_USEC_TO_DATETIME * int64(1000)
	UA_SEC_TO_DATETIME  = UA_MSEC_TO_DATETIME * int64(1000)
)

/* Datetime of 1 Jan 1970 00:00 UTC */
const UA_DATETIME_UNIX_EPOCH = int64(11644473600) * UA_SEC_TO_DATETIME

/* The current time */
func UA_DateTime_now() UA_DateTime {
	return UA_DateTime(C.UA_DateTime_now())
}

/* CPU clock invariant to system time changes. Use only for time diffs, not current time */
func UA_DateTime_nowMonotonic(void) UA_DateTime {
	return UA_DateTime(C.UA_DateTime_nowMonotonic())
}

type UA_DateTimeStruct C.UA_DateTimeStruct

func UA_DateTime_toStruct(dt UA_DateTime) UA_DateTimeStruct {
	return UA_DateTimeStruct(C.UA_DateTime_toStruct(dt))
}

func UA_DateTime_toString(dt UA_DateTime) UA_String {
	return UA_DateTime_toString(C.UA_DateTime_toStruct(dt))
}

/**************************************************************************/
/* Guid: A 16 byte value that can be used as a globally unique identifier */
/**************************************************************************/
type UA_Guid C.UA_Guid

type UA_ByteString C.UA_ByteString
type UA_XmlElement C.UA_XmlElement
type UA_NodeId C.UA_NodeId
type UA_ExpandedNodeId C.UA_ExpandedNodeId
type UA_QualifiedName C.UA_QualifiedName
type UA_LocalizedText C.UA_LocalizedText
type UA_ExtensionObject C.UA_ExtensionObject
type UA_Variant C.UA_Variant
type UA_DataValue C.UA_DataValue
type UA_DiagnosticInfo C.UA_DiagnosticInfo
