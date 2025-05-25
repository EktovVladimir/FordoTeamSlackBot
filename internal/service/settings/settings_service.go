package settings

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"

type settingsService struct {
	dbRepository abstract.DbRepository
}

func New(dbRepository abstract.DbRepository) *settingsService {
	return &settingsService{dbRepository}
}

func (s *settingsService) GetStringItem(source abstract.Source, key string) string {
	//TODO логика поиска настройки c учётом wildcard
	item, err := s.dbRepository.GetItem(source.GetKey(), key)

	if err != nil {
		return ""
	}

	return item.Value
}
