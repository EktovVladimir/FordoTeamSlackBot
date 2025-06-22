package main

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/pkg/grpc/proto/user_finder"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	environment.InitGlobal()
	cfg := config.Load("team_frodo")
	addr := fmt.Sprintf("%s:%d", cfg.GrpcServer.Host, cfg.GrpcServer.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := user_finder_grpc.NewUserFinderServiceClient(conn)

	res, err := client.FindByEmailWithDb(ctx, &user_finder_grpc.FindByEmailRequest{Email: "ektov@mego.travel"})
	if err != nil {
		log.Fatalf("Failed on FindByEmailWithDb: %v", err)
	}
	log.Println("FindByEmailWithDb result", res)

	res2, err := client.FindByEmail(ctx, &user_finder_grpc.FindByEmailRequest{Email: "ektov@mego.travel"})
	if err != nil {
		log.Fatalf("Failed on FindByEmail: %v", err)
	}
	log.Println("FindByEmail result", res2)

	res3, err := client.SaveUser(ctx, &user_finder_grpc.SaveUserRequest{User: res2.User})
	if err != nil {
		log.Fatalf("Failed on SaveUser: %v", err)
	}
	log.Println("SaveUser result", res3)
}
