package lyncrpc

type Bot interface {
	Id() string
	URI() string
	Shutdown() error
	Login() error
	Relogin() error
	Bind() error
	Online() error
	Away() error
	User() string
	RemoteURI() string
	Flush() int64
	IsLoggedOut() bool
	Error() chan error
}
