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

type loginRequest struct{}
type loginResponse struct{}

func (client *client) Login() error {
	return client.Call("LOGIN", &loginRequest{}, &loginResponse{})
}
