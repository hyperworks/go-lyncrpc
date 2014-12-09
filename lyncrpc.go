package lyncrpc

type LyncRPC interface{
	Login() error
}

type loginRequest struct{}
type loginResponse struct{}
