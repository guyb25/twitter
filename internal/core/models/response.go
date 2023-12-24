package models

type Response struct {
	Status  int
	Message interface{}
}

func NewResponse(status int, message interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}
