package errors

// Error an internal error use for represent
type Error struct {
	Message    string `json:"message"`
	httpStatus int
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) HttpStatus() int {
	return e.httpStatus
}

func New(msg string, status int) *Error {
	return &Error{
		Message:    msg,
		httpStatus: status,
	}
}
