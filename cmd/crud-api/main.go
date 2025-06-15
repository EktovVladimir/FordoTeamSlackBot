package main

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/crud-api"

func main() {
	crud_api.Run().Wait()
}
