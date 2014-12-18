package lyncrpc

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type client struct {
	*rpc.Client
}

func Dial(network, address string) (Interface, error) {
	rpcClient, e := jsonrpc.Dial(network, address)
	if e != nil {
		return nil, e
	}

	return &client{rpcClient}, nil
}

func (client *client) Hello(name string) (string, error) {
	response := ""
	e := client.Call("HELLO", arg("Name", name), &response)
	return response, e
}

func (client *client) Ping() error {
	response := ""
	e := client.Call("PING", nil, &response)
	if e == nil && response != "PONG" {
		e = fmt.Errorf("PONG reply not received!")
	}
	return e
}

func (client *client) Date() (time.Time, error) {
	response := ""
	e := client.Call("DATE", map[string]string{}, &response)
	if e != nil {
		return time.Time{}, nil
	}

	return time.Parse(time.RFC3339, response)
}

func (client *client) SignIn(serverUrl, address, username, password string) error {
	request := map[string]string{
		"ServerUrl":     serverUrl,
		"SignInAddress": address,
		"Username":      username,
		"Password":      password,
	}
	return client.Call("SIGNIN", request, nil)
}

func (client *client) SignOut() error {
	return client.Call("SIGNOUT", nil, nil)
}

func (client *client) Availability() (string, error) {
	response := ""
	e := client.Call("AVAILABILITY", nil, &response)
	return response, e
}

func (client *client) SetAvailability(availability string) error {
	return client.Call("SET_AVAILABILITY", arg("Availability", availability), nil)
}

func (client *client) Contacts() ([]*Contact, error) {
	result := []*Contact{}
	e := client.Call("CONTACTS", nil, &result)
	return result, e
}

func (client *client) BeginConversation(recipientUri string) error {
	return client.Call("BEGIN_CONVERSATION", arg("RecipientUri", recipientUri), nil)
}

func (client *client) AcceptConversation() error {
	return client.Call("ACCEPT_CONVERSATION", nil, nil)
}

func (client *client) SendMessage(message string) error {
	return client.Call("SEND_MESSAGE", arg("Message", message), nil)
}

func (client *client) ReceiveMessage() (string, error) {
	response := ""
	e := client.Call("RECEIVE_MESSAGE", nil, &response)
	return response, e
}

func (client *client) EndConversation() error {
	return client.Call("END_CONVERSATION", nil, nil)
}

func arg(key, value string) map[string]string {
	return map[string]string{key: value}
}
