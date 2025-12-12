package replies

type ErrorCode int

const (
	ErrUnknownCommand ErrorCode = iota
	ErrInvalidArgs
	ErrSyntaxError
)

type AppError struct {
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	UnknownCommand = &AppError{
		Code:    ErrUnknownCommand,
		Message: "ERR: unknown command",
	}

	InvalidArgs = &AppError{
		Code:    ErrInvalidArgs,
		Message: "ERR: invalid arguments",
	}

	SyntaxError = &AppError{
		Code:    ErrSyntaxError,
		Message: "ERR: syntax error",
	}
)
