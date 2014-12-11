package main

import (
	"fmt"
	"github.com/hyperworks/go-lyncrpc"
)

func main() {
	client, e := lyncrpc.Dial("tcp", "192.168.1.35:7331")
	noError(e)

	var server, username, password string
	fmt.Printf("server: ")
	_, e = fmt.Scanln(&server)
	noError(e)
	fmt.Printf("username: ")
	_, e = fmt.Scanln(&username)
	noError(e)
	fmt.Printf("password: ")
	_, e = fmt.Scanln(&password)
	noError(e)

	e = client.SignIn(server, username, password)
	noError(e)
}

func noError(e error) {
	if e != nil {
		panic(e)
	}
}
