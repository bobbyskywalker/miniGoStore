package errors

type ErrorCode int

const (
	ErrUnknownCommand ErrorCode = iota
	ErrInvalidArgs
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
)
