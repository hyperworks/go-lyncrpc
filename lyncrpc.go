package lyncrpc

type LyncRPC interface {
	Hello(name string) (string, error)
	Login() error
}
