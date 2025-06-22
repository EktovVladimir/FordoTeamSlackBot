package user_finder

import (
	"context"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/mapping"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils"
	"github.com/google/go-github/v72/github"
	"github.com/slack-go/slack"
)

var (
	ErrorSlackEmailNotPresented = errors.New("email not presented by slack")
	ErrorGithubUserNotFound     = errors.New("github user not found")
)

type UserFinder struct {
	userRepo     repository.UserRepository
	slackClient  *slack.Client
	githubClient *github.Client
}

func New(userRepo repository.UserRepository,
	slackClient *slack.Client,
	githubClient *github.Client) *UserFinder {
	return &UserFinder{userRepo: userRepo, slackClient: slackClient, githubClient: githubClient}
}

func (u *UserFinder) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	res := &models.User{}

	slackUser, err := u.slackClient.GetUserByEmailContext(ctx, email)
	if err != nil {
		return nil, err
	}

	res.Email = email
	res.SlackName = slackUser.Name
	res.SlackId = slackUser.ID

	ghUsers, _, err := u.githubClient.Search.Users(ctx, fmt.Sprintf("%s in:email", email), &github.SearchOptions{})
	if err != nil {
		return res, err
	}

	if ghUsers.Total == nil || *ghUsers.Total == 0 {
		return res, ErrorGithubUserNotFound
	}

	res.GithubLogin = utils.DerefString(ghUsers.Users[0].Login)

	return res, nil
}

func (u *UserFinder) FindUserByEmailWithDb(ctx context.Context, email string) (*models.User, error) {
	dbUser, err := u.userRepo.GetByEmail(ctx, email)
	if err == nil && dbUser.IsFullFilled() {
		return mapping.MapDbUserToModel(dbUser), nil
	}

	return u.FindUserByEmail(ctx, email)
}

func (u *UserFinder) SaveUser(ctx context.Context, user *models.User) error {
	var dbUser *db.User
	if user.Id > 0 {
		dbUser, _ = u.userRepo.GetById(ctx, user.Id)
	} else if user.Email != "" {
		dbUser, _ = u.userRepo.GetByEmail(ctx, user.Email)
	} else {
		return errors.New("user must have email or id")
	}

	mapped := mapping.MapUserToDb(user)

	if dbUser == nil {
		if err := u.userRepo.Create(ctx, mapped); err != nil {
			return err
		}
	} else {
		mapped.Id = dbUser.Id
		if err := u.userRepo.Update(ctx, mapped); err != nil {
			return err
		}
	}

	return nil
}
