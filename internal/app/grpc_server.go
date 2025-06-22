package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/pkg/grpc/proto/user_finder"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	app              *app
	userFinderServer user_finder_grpc.UserFinderServiceServer
}

func newGrpcServer(app *app, userFinderServer user_finder_grpc.UserFinderServiceServer) *grpcServer {
	return &grpcServer{app: app, userFinderServer: userFinderServer}
}

func (s *grpcServer) Start(ctx context.Context) {
	serverConf := s.app.cfg.GrpcServer
	addr := fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	gs := grpc.NewServer()

	user_finder_grpc.RegisterUserFinderServiceServer(gs, s.userFinderServer)

	reflection.Register(gs)

	serverErr := make(chan error, 1)
	go func() {
		logrus.Infof("%s grpcServer starting on %s", s.app.appName, addr)
		if err = gs.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			serverErr <- fmt.Errorf("failed to serve grpcServer: %w", err)
		}
	}()

	select {
	case err = <-serverErr:
		logrus.Errorf("%s grpcServer has error: %v", s.app.appName, err)
	case <-ctx.Done():
		logrus.Infof("%s grpcServer received shutdown signal", s.app.appName)
	}

	gs.GracefulStop()

	logrus.Infof("%s grpcServer stopped gracefully", s.app.appName)

}
