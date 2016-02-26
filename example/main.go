/*
  To run this example, start the example server from open62541
  and execute:
    go run main.go
*/

package main

import (
	"fmt"

	"github.com/coussej/gopcua"
)

func main() {
	c := gopcua.New_UA_Client(gopcua.UA_ClientConfig_standard)
	_ = c.Connect("opc.tcp://localhost:16664")

	ni := gopcua.New_UA_NodeId_String(1, "the.answer")

	rvi := gopcua.UA_ReadValueId{
		NodeId:      ni,
		AttributeId: uint32(gopcua.UA_ATTRIBUTEID_VALUE),
	}

	r := gopcua.UA_ReadRequest{}

	//r.Init()

	r.NodesToRead = &rvi
	r.NodesToReadSize = 1

	resp := c.Read(r)

	//result, err := c.Read_Int32("the.answer")
	fmt.Printf("the.answer: %+v\n", r.NodesToRead)
	fmt.Printf("response: %+v\n", *resp.Results.Value.Type)
	fmt.Printf("response: %+v\n", gopcua.UA_Int32(*resp.Results.Value.Data))
}

func stringToByteArray(s string) [16]byte {
	var arr [16]byte
	for k, v := range []byte(s) {
		arr[k] = byte(v)
	}
	return arr
}
