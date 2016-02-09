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
	c := gopcua.NewClient()
	err := c.Connect("opc.tcp://localhost:16664")

	if err != nil {
		panic(err)
	}

	result, err := c.Read_Int32("the.answer")
	fmt.Printf("the.answer: %+v\n", result)
}
