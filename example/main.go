package main

import (
	"fmt"
	"time"
	"github.com/hyperworks/go-lyncrpc"
)

func main() {
	client, e := lyncrpc.Dial("tcp", "0.0.0.0:7331")
	noError(e)
	hello, e := client.Hello("chakrit")
	noError(e)

	fmt.Println(hello)
	for {
		t, e := client.Date()
		noError(e)

		fmt.Println(t.Format(time.RFC822Z))
	}
}

func noError(e error) {
	if e != nil {
		panic(e)
	}
}
