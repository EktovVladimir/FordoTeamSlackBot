package main

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/hw"
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

	hw.RunHw16(ctx).Wait()

	fmt.Println("Application finished")
}
