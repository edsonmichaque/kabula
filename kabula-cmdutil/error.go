package libcmd

type Error struct {
	ExitCode int
	Err      error
}

func NewError(code int, err error) Error {
	return Error{
		ExitCode: code,
		Err:      err,
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
