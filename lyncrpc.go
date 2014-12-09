package lyncrpc

type LyncRPC interface {
	Hello(name string) (string, error)
	Login() error
}

type helloRequest struct{ name string }
type helloResponse struct{ line string }

type loginRequest struct{}
type loginResponse struct{}
