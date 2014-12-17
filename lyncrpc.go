package lyncrpc

import (
	"time"
)

type Contact struct {
	Uri    string
	Name   string
	Groups []string
}

type LyncRPC interface {
	Hello(name string) (string, error)
	Date() (time.Time, error)
	SignIn(serverUrl, address, username, password string) error
	SignOut() error
	Availability() (string, error)
	SetAvailability(string) error
	Contacts() ([]*Contact, error)
}
