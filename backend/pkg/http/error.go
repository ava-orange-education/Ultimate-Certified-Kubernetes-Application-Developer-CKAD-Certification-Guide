package http

type Error struct {
	HTTPStatus int
	ErrCode    string
	Message    string
}

func (e Error) Error() string {
	var str string

	if e.ErrCode != "" {
		str = e.ErrCode
	}

	if e.Message != "" {
		str += ": " + e.Message
	}

	return str
}

func NewError(status int, code, message string) error {
	return &Error{
		HTTPStatus: status,
		ErrCode:    code,
		Message:    message,
	}
}

func (e Error) ToAPIResponse() ErrorResponse {
	return ErrorResponse{
		Code:    e.ErrCode,
		Message: e.Message,
	}
}
