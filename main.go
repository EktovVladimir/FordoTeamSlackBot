package main

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/run"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	run.RunHw16(ctx).Wait()

	fmt.Println("Application finished")
}
