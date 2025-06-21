package settings_handler

// CreateSettingRequest model
//
//	@Description	Запрос на создание настройки
type CreateSettingRequest struct {
	//	@Description	Уникальный ключ настройки
	//	@Required		true
	Key string `json:"key" validate:"required,min=1"`

	//	@Description	Значение настройки
	//	@Example:		10485760
	//	@Required		true
	Value string `json:"value" validate:"required,min=1"`
}

// UpdateSettingRequest model
//
//	@Description	Запрос на обновление настройки
type UpdateSettingRequest struct {
	//	@Description	Уникальный ключ настройки
	//	@Required		true
	Key string `json:"key" validate:"required,min=1"`

	//	@Description	Новое значение настройки
	//	@Required		true
	Value string `json:"value" validate:"required,min=1"`
}
