package main

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/management_api"

func main() {
	management_api.Run().Wait()
}
