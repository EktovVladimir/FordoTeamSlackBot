package main

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/run"
)

func main() {
	run.RunHw16().Wait()

	fmt.Println("Application finished")
}
