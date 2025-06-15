package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomStringFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func RandomPRNumber() string {
	return fmt.Sprintf("%d", rand.Intn(1000)+1)
}

func RandomRepo() string {
	repos := []string{
		"FordoTeamSlackBot",
		"GolangPracticeOtus",
	}
	return RandomStringFromList(repos)
}

func RandomBranch() string {
	branches := []string{
		"feature/test1",
		"feature/test3",
		"bugfix/2025-05-30",
	}
	return RandomStringFromList(branches)
}

func RandomAction() string {
	actions := []string{
		"opened",
		"closed",
		"reopened",
		"merged",
		"synchronize",
	}
	return RandomStringFromList(actions)
}

func RandomRequester() string {
	requesters := []string{
		"EktovVladimir",
		"developer1",
		"coder42",
		"git-user",
	}
	return RandomStringFromList(requesters)
}

func RandomChannel() string {
	channels := []string{
		"general",
		"random",
		"devops",
		"backend",
		"frontend",
		"releases",
	}
	return "#" + RandomStringFromList(channels)
}

func RandomUserName() string {
	users := []string{
		"alice",
		"bob",
		"charlie",
		"dave",
		"eve",
		"frank",
	}
	return RandomStringFromList(users)
}

func RandomCommandName() string {
	users := []string{
		"cr",
		"deploy",
	}
	return RandomStringFromList(users)
}

func RandomCommandArgs() string {
	users := []string{
		"new",
		fmt.Sprintf("pr %s", RandomPRNumber()),
	}
	return RandomStringFromList(users)
}

func GetCurrentThreadTS() string {
	now := time.Now()
	// Unix() возвращает секунды, UnixNano() наносекунды (1e9)
	seconds := now.Unix()
	nanos := now.UnixNano() % 1e9 // Остаток от деления на 1e9 (наносекунды)
	micros := nanos / 1e3         // Переводим наносекунды в микросекунды
	return fmt.Sprintf("%d.%06d", seconds, micros)
}
