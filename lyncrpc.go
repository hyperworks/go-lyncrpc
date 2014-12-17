package lyncrpc

import (
	"time"
)

type Contact struct {
	Uri    string
	Name   string
	Groups []string
}

// TODO: Probably better to split this up.
type Interface interface {
	Hello(name string) (string, error)
	Ping() error
	Date() (time.Time, error)
	SignIn(serverUrl, address, username, password string) error
	SignOut() error
	Availability() (string, error)
	SetAvailability(availability string) error
	Contacts() ([]*Contact, error)
	BeginConversation(recipientUri string) error
	SendMessage(message string) error
	EndConversation() error
}
