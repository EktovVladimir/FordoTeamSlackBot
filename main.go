package main

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/run"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	run.RunHw16(ctx).Wait()

	fmt.Println("Application finished")
}
