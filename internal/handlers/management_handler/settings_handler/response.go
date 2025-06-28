package settings_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

// SettingResponse model
// @Description Настройка ключа-значение
type SettingResponse struct {
	// @Description Уникальный Id
	Id db.UniqId `json:"id"`
	// @Description Уникальный ключ настройки
	Key string `json:"key"`
	// @Description Значение настройки
	Value string `json:"value"`
}
