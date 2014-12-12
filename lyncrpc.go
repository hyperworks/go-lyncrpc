package lyncrpc

import (
	"time"
)

type LyncRPC interface {
	Hello(name string) (string, error)
	Date() (time.Time, error)
	SignIn(serverUrl, username, password string) error
	SignOut() error
}
