package main

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service"
	"time"
)

func main() {
	for i := 0; i < 25; i++ {
		service.GenerateAndStore()

		time.Sleep(100 * time.Millisecond)
	}
}
