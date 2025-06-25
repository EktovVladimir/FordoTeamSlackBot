package review_manager

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/gh_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/jira_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/slack_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"golang.org/x/net/context"
	"strconv"
)

type source interface {
	GetChannelId() string
	GetEmail() string
}

type ReviewManager struct {
	repo            repository.CodeReviewRepository
	userFinder      *user_finder.UserFinder
	commitRetriever *gh_service.CommitRetriever
	prService       *gh_service.PullRequestService
	jrService       *jira_service.IssueService
	messagePoster   *slack_service.MessagePoster
}

func NewReviewManager(
	repo repository.CodeReviewRepository,
	userFinder *user_finder.UserFinder,
	commitRetriever *gh_service.CommitRetriever,
	prService *gh_service.PullRequestService,
	jrService *jira_service.IssueService,
	messagePoster *slack_service.MessagePoster) *ReviewManager {
	return &ReviewManager{
		repo,
		userFinder,
		commitRetriever,
		prService,
		jrService,
		messagePoster}
}

func (r *ReviewManager) RequestReview(
	ctx context.Context,
	src source,
	prRefs []*models.PullRequestRef,
	opts ...RequestReviewOption) (res *models.CodeReviewFact, err error) {

	conf := defaultCreateThreadConfig()
	if err = conf.setOptions(opts...); err != nil {
		return nil, err
	}

	email := src.GetEmail()

	reqUser, err := r.userFinder.FindUserByEmailWithDb(ctx, email)
	if err != nil {
		return nil, err
	}

	ghReviewers, err := r.prService.GetReviewersMany(ctx, prRefs)
	if err != nil {
		return nil, err
	}

	revUsers, err := r.userFinder.FindGhReviewerMany(ctx, ghReviewers)
	if err != nil {
		return nil, err
	}

	issuePattern := fmt.Sprintf(`(?i)(\s|^)%s-[0-9]+`, conf.jiraProject)
	issueKeys, err := r.commitRetriever.GetIssueNumbersMany(ctx, prRefs, issuePattern)

	issues, err := r.jrService.GetIssueList(ctx, issueKeys)

	createdMsg, err := r.messagePoster.CreateReviewThread(ctx, src, &models.CrTreadMessage{
		Requester: reqUser,
		Reviewers: revUsers,
		Issues:    issues,
		Prs:       prRefs,
	})
	if err != nil {
		return nil, err
	}

	res = &models.CodeReviewFact{
		Requester: reqUser,
		Reviewers: revUsers,
		Prs:       prRefs,
		Issues:    issues,
		Message:   createdMsg,
	}

	dbModel := &db.CodeReview{
		ThreadTs: createdMsg.Ts,
		//Пока достаточно ссылки на 1 pr, в будущем сделать отдельную связь
		PullRequestNumber: strconv.Itoa(prRefs[0].Number),
		Status:            "pending",
	}

	err = r.repo.Create(ctx, dbModel)
	if err != nil {
		return nil, err
	}

	res.Id = dbModel.Id

	return res, nil
}
