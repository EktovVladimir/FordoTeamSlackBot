package api

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CreateSettingRequest struct {
	Key   string `json:"key" validate:"required,min=1"`
	Value string `json:"value" validate:"required,min=1"`
}

type UpdateSettingRequest struct {
	Key   string `json:"key" validate:"required,min=1"`
	Value string `json:"value" validate:"required,min=1"`
}

type SettingResponse struct {
	Id    types.UniqId `json:"id"`
	Key   string       `json:"key"`
	Value string       `json:"value"`
}
