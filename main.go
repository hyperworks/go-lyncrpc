package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type TestService struct{}

type HelloArg struct{ Name string }
type HelloResult struct{ Line string }

func (ts *TestService) Hello(args *HelloArg, result *HelloResult) error {
	*result = HelloResult{
		Line: "hello, " + args.Name,
	}
	return nil
}

func main() {
	switch os.Args[1] {
	case "client":
		startClient()
	case "server":
		startServer()
	default:
		log.Fatal("specify either `server` or `client` mode.")
	}
}

func startClient() {
	client, e := jsonrpc.Dial("tcp", "0.0.0.0:7331")
	noError(e)

	arg := &HelloArg{"Chakrit"}
	line := &HelloResult{}
	e = client.Call("TestService.Hello", arg, line)
	noError(e)

	log.Println(line.Line)
}

func startServer() {
	log.Println("starting server...")
	server := rpc.NewServer()
	server.Register(&TestService{})

	listener, e := net.Listen("tcp", "0.0.0.0:7331")
	noError(e)

	for {
		conn, e := listener.Accept()
		noError(e)

		log.Println("processing request...")
		go func() {
			server.ServeCodec(jsonrpc.NewServerCodec(conn))
			log.Println("done.")
		}()
	}
}

func noError(e error) {
	if e != nil {
		log.Println("wait...")
	}
}
