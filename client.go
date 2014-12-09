package lyncrpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type client struct {
	*rpc.Client
}

func Dial(network, address string) (LyncRPC, error) {
	rpcClient, e := jsonrpc.Dial(network, address)
	if e != nil {
		return nil, e
	}

	return &client{rpcClient}, nil
}

func (client *client) Hello(name string) (string, error) {
	request := &struct{ Name string }{name}
	response := ""
	e := client.Call("HELLO", request, &response)
	return response, e
}

func (client *client) Login() error {
	return client.Call("LOGIN", nil, nil)
}
