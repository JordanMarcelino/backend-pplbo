package models

type Info struct {
	Success bool   `json:"success"`
	Meta    any    `json:"meta"`
	Message string `json:"message"`
}

type WebResponse[T any] struct {
	Info Info `json:"info"`
	Data T    `json:"data"`
}

func NewErrorResponse(message string) WebResponse[*string] {
	return WebResponse[*string]{
		Info: Info{
			Success: false,
			Meta:    nil,
			Message: message,
		},
		Data: nil,
	}
}

func NewSuccessResponse[T any](data T, message string) WebResponse[T] {
	return WebResponse[T]{
		Info: Info{
			Success: true,
			Meta:    nil,
			Message: message,
		},
		Data: data,
	}
}
