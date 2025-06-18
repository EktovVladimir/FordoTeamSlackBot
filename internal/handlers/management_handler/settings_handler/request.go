package settings_handler

type CreateSettingRequest struct {
	Key   string `json:"key" validate:"required,min=1"`
	Value string `json:"value" validate:"required,min=1"`
}

type UpdateSettingRequest struct {
	Key   string `json:"key" validate:"required,min=1"`
	Value string `json:"value" validate:"required,min=1"`
}
