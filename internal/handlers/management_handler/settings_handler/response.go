package settings_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type SettingResponse struct {
	Id    types.UniqId `json:"id"`
	Key   string       `json:"key"`
	Value string       `json:"value"`
}
