/*
  Copyright 2016 Jeroen Coussement.
*/

package gopcua

/*
#cgo CFLAGS: -std=c99 -I .
#include <stdlib.h>
#include <gopcua_client.h>
#include <open62541.h>
*/
import "C"

import "unsafe"

func NewClient() Client {
	c := C.gopcua_client_new()
	return Client{c}
}

func (c *Client) Connect(ConnInfo string) (err error) {
	ci := C.CString(ConnInfo)
	defer C.free(unsafe.Pointer(ci))
	_, err = C.gopcua_client_connect(c.c_client, ci)
	return
}

func (c *Client) Read_Int32(NodeId string) (p Point_Int32, err error) {
	n := C.CString(NodeId)
	defer C.free(unsafe.Pointer(n))
	p_c, err := C.gopcua_read_int32(c.c_client, n)
	if err == nil {
		p = p_c.toPoint()
	}
	return
}

func (c *Client) Read_Float32(NodeId string) (p Point_Float32, err error) {
	n := C.CString(NodeId)
	defer C.free(unsafe.Pointer(n))
	p_c, err := C.gopcua_read_float32(c.c_client, n)
	if err == nil {
		p = p_c.toPoint()
	}
	return
}
