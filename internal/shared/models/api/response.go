package api

type Response[T any] struct {
	Data  T      `json:"data"`
	Error *Error `json:"error,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) Response[any] {
	return Response[any]{
		Error: &Error{
			Message: message,
		},
	}
}
