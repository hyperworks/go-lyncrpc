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

/*
	Buddies() ([]Buddy, error)
	AddBuddy(jid, name, group string) error
	DeleteBuddy(jid string) error

	Block(uri ...string) error
	Unblock(uri ...string) error

	CheckPresence(id string) (presence.Activity, error)
	SetPresence(pres presence.Activity) error
	Subscribe(uri string) error
	Unsubscribe(uri string) error
	ApproveSubscribe(uri string) error

	ChatSession(uri string) (ChatSession, error)
	AcceptChatSession(uri string) (ChatSession, error)
	JoinRoom(roomName, nick string) (GroupChatSession, bool, error)
	RequestProvisioning(prefix string) error
	*/
