package settings

import (
	abstract2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/abstract"
)

type settingsService struct {
	dbRepository abstract2.DbRepository
}

func New(dbRepository abstract2.DbRepository) *settingsService {
	return &settingsService{dbRepository}
}

func (s *settingsService) GetStringItem(source abstract2.Source, key string) string {
	//TODO логика поиска настройки c учётом wildcard
	item, err := s.dbRepository.GetItem(source.GetKey(), key)

	if err != nil {
		return ""
	}

	return item.Value
}
