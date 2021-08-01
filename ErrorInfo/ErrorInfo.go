package ErrorInfo

type ConnectError struct {
	ErrCode int
	ErrMsg string
}

func (err *ConnectError) Error()string{
	return err.ErrMsg
}

var DbConnectionError = &ConnectError{10001,"Failed connection error."}
