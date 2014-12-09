package lyncrpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Client struct {
	rpc *rpc.Client
}

func Dial(network, address string) (LyncRPC, error) {
	rpcClient, e := jsonrpc.Dial(network, address)
	if e != nil {
		return nil, e
	}

	return &Client{rpc: rpcClient}, nil
}

func (client *Client) Login() error {
	return client.rpc.Call("LOGIN", &loginRequest{}, &loginResponse{})
}
