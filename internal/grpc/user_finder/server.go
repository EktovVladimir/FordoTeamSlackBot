package user_finder_server

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/pkg/grpc/proto/user_finder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	user_finder_grpc.UnimplementedUserFinderServiceServer
	service *user_finder.UserFinder
}

func (s Server) FindByEmail(ctx context.Context, request *user_finder_grpc.FindByEmailRequest) (*user_finder_grpc.UserResponse, error) {
	user, err := s.service.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return &user_finder_grpc.UserResponse{
			Error: err.Error(),
		}, status.Error(codes.NotFound, err.Error())
	}

	return &user_finder_grpc.UserResponse{
		User: convertModelToProto(user),
	}, nil
}

func (s Server) FindByEmailWithDb(ctx context.Context, request *user_finder_grpc.FindByEmailRequest) (*user_finder_grpc.UserResponse, error) {
	user, err := s.service.FindUserByEmailWithDb(ctx, request.Email)
	if err != nil {
		return &user_finder_grpc.UserResponse{
			Error: err.Error(),
		}, status.Error(codes.NotFound, err.Error())
	}

	return &user_finder_grpc.UserResponse{
		User: convertModelToProto(user),
	}, nil
}

func (s Server) SaveUser(ctx context.Context, request *user_finder_grpc.SaveUserRequest) (*user_finder_grpc.SaveUserResponse, error) {
	err := s.service.SaveUser(ctx, convertProtoToModel(request.User))
	if err != nil {
		return &user_finder_grpc.SaveUserResponse{
			Error: err.Error(),
		}, status.Error(codes.Internal, err.Error())
	}

	return &user_finder_grpc.SaveUserResponse{
		Success: true,
	}, nil
}

func New(userFinder *user_finder.UserFinder) *Server {
	return &Server{service: userFinder}
}

func convertModelToProto(user *models.User) *user_finder_grpc.User {
	return &user_finder_grpc.User{
		Id:          uint64(user.Id),
		Email:       user.Email,
		SlackName:   user.SlackName,
		SlackId:     user.SlackId,
		GithubLogin: user.GithubLogin,
	}
}

func convertProtoToModel(user *user_finder_grpc.User) *models.User {
	return &models.User{
		Id:          db.UniqId(user.Id),
		Email:       user.Email,
		SlackName:   user.SlackName,
		SlackId:     user.SlackId,
		GithubLogin: user.GithubLogin,
	}
}
