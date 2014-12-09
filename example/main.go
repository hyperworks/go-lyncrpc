package main

import (
	"fmt"
	"github.com/hyperworks/go-lyncrpc"
)

func main() {
	client, e := lyncrpc.Dial("tcp", "0.0.0.0:7331")
	noError(e)
	line, e := client.Hello("chakrit")
	noError(e)

	fmt.Println(line)
}

func noError(e error) {
	if e != nil {
		panic(e)
	}
}
