package lyncrpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
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

func (client *client) Date() (time.Time, error) {
	response := ""
	e := client.Call("DATE", map[string]string{}, &response)
	if e != nil {
		return time.Time{}, nil
	}

	return time.Parse(time.RFC3339, response)
}

func (client *client) SignIn(serverUrl, username, password string) error {
	request := map[string]string{
		"ServerUrl": serverUrl,
		"Username":  username,
		"Password":  password,
	}
	return client.Call("SIGNIN", request, nil)
}

func (client *client) Login() error {
	return client.Call("LOGIN", nil, nil)
}
