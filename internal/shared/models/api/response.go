package api

// Response base model
// @Description Базовая модель ответа на запрос
type Response[T any] struct {
	Data  T      `json:"data"`            //Данные
	Error *Error `json:"error,omitempty"` //Ошибка
}

// Error model
// @Description Модель ошибки на запрос
type Error struct {
	Message string `json:"message"` //Текст ошибки
}

func NewErrorResponse(message string) Response[any] {
	return Response[any]{
		Error: &Error{
			Message: message,
		},
	}
}
