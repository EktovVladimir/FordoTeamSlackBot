package mapping

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

func MapDbUserToModel(in *db.User) *models.User {
	return &models.User{
		Id:          in.Id,
		Email:       in.Email,
		SlackId:     in.SlackId,
		SlackName:   in.SlackName,
		GithubLogin: in.GithubName,
	}
}

func MapUserToDb(in *models.User) *db.User {
	return &db.User{
		UniqFields: db.WithId(in.Id),
		Email:      in.Email,
		SlackId:    in.SlackId,
		SlackName:  in.SlackName,
		GithubName: in.GithubLogin,
	}
}
