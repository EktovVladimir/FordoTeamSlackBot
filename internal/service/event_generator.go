package service

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"math/rand"
)

func GenerateEvent() abstract.Event {
	//Для выполнения ДЗ, это генератор случайных событий,
	//В будущем планируются реальные слушатели команд бота,
	//вызовы из api и веб-хуки github

	eventType := rand.Intn(2)

	switch eventType {
	case 0:
		return event.GenerateRandomSlackCommandEvent()
	default:
		return event.GenerateRandomPullRequestEvent()
	}
}

func GenerateAndStore() {
	event := GenerateEvent()
	repository.StoreEvent(event)
}
