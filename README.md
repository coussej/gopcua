# gopcua

This package wants to provide a thin, high-level wrapper around [open62541](https://github.com/open62541/open62541), so that you can have access to some basic OPC UA client functionality from within Go. 

> NOTE: still very experimental and not properly tested.

## Usage

See example. To run it, start the [example server](https://github.com/open62541/open62541/blob/master/examples/server.c) from open62541 and execute:
```    
go run main.go
```

```go
package main

import (
	"fmt"

	"github.com/coussej/gopcua"
)

func main() {
  // get a client and connect to the server
	c := gopcua.NewClient()
	err := c.Connect("opc.tcp://localhost:16664")

	if err != nil {
		panic(err)
	}

	result, err := c.Read_Int32("the.answer")
	fmt.Printf("the.answer: %+v\n", result)
}
```

