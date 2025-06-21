package settings_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

// SettingResponse model
// @Description Настройка ключа-значение
type SettingResponse struct {
	// @Description Уникальный Id
	Id types.UniqId `json:"id"`
	// @Description Уникальный ключ настройки
	Key string `json:"key"`
	// @Description Значение настройки
	Value string `json:"value"`
}
