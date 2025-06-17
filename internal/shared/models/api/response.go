package api

type Response[T any] struct {
	Data  T      `json:"data"`
	Error *Error `json:"error,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}
