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

var UA_ClientConfig_standard = UA_ClientConfig{
	Timeout:               5000,
	SecureChannelLifeTime: 600000,
	LocalConnectionConfig: UA_ConnectionConfig{
		ProtocolVersion: 0,
		SendBufferSize:  65536,
		RecvBufferSize:  65536,
		MaxMessageSize:  65536,
		MaxChunkCount:   1,
	},
}

func New_UA_Client(config UA_ClientConfig) *UA_Client {
	cnf := *(*C.UA_ClientConfig)(unsafe.Pointer(&config))
	clt := C.gopcua_client_new(cnf)
	return (*UA_Client)(unsafe.Pointer(clt))
}

func (c *UA_Client) Reset() {
	clt := (*C.UA_Client)(unsafe.Pointer(c))
	C.UA_Client_reset(clt)
}

func (c *UA_Client) Delete() {
	clt := (*C.UA_Client)(unsafe.Pointer(c))
	C.UA_Client_delete(clt)
}

func (c *UA_Client) Connect(ConnInfo string) (s UA_StatusCode) {
	ci := C.CString(ConnInfo)
	defer C.free(unsafe.Pointer(ci))

	clt := (*C.UA_Client)(unsafe.Pointer(c))
	stat := C.gopcua_client_connect(clt, ci)
	s = *(*UA_StatusCode)(unsafe.Pointer(&stat))
	return
}

func (c *UA_Client) Read(r UA_ReadRequest) UA_ReadResponse {
	clt := (*C.UA_Client)(unsafe.Pointer(c))
	req := *(*C.UA_ReadRequest)(unsafe.Pointer(&r))
	resp := C.UA_Client_Service_read(clt, req)
	return *(*UA_ReadResponse)(unsafe.Pointer(&resp))
}

// Initialize Node ID's

func New_UA_NodeId_Numeric(namespace UA_UInt16, id UA_UInt32) UA_NodeId {
	cs := *(*C.UA_UInt32)(unsafe.Pointer(&id))
	ns := *(*C.UA_UInt16)(unsafe.Pointer(&namespace))

	node := C.UA_NODEID_NUMERIC(ns, cs)
	return *(*UA_NodeId)(unsafe.Pointer(&node))
}

func New_UA_NodeId_String(namespace UA_UInt16, id string) UA_NodeId {
	cs := C.CString(id)
	defer C.free(unsafe.Pointer(cs))
	ns := *(*C.UA_UInt16)(unsafe.Pointer(&namespace))

	node := C.UA_NODEID_STRING(ns, cs)
	return *(*UA_NodeId)(unsafe.Pointer(&node))
}

func New_UA_NodeId_String_Alloc(namespace UA_UInt16, id string) UA_NodeId {
	cs := C.CString(id)
	defer C.free(unsafe.Pointer(cs))
	ns := *(*C.UA_UInt16)(unsafe.Pointer(&namespace))

	node := C.UA_NODEID_STRING_ALLOC(ns, cs)
	return *(*UA_NodeId)(unsafe.Pointer(&node))
}

func New_UA_NodeId_Guid(namespace UA_UInt16, id UA_Guid) UA_NodeId {
	cs := *(*C.UA_Guid)(unsafe.Pointer(&id))
	ns := *(*C.UA_UInt16)(unsafe.Pointer(&namespace))

	node := C.UA_NODEID_GUID(ns, cs)
	return *(*UA_NodeId)(unsafe.Pointer(&node))
}
