package github

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	git2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"math/rand"
	"strconv"
)

type pullRequestService struct {
	config config.Github
	//TODO клиент GH
}

func New(config config.Github) *pullRequestService {
	//TODO создание клиента GH
	return &pullRequestService{
		config: config,
	}
}

func (p *pullRequestService) Get(repo string, number string) (git2.PullRequest, error) {
	//TODO implement me
	res := git2.PullRequest{
		Number:     number,
		Repo:       repo,
		Branch:     "develop",
		BaseBranch: "master",
		State:      "opened",
		Requester:  git2.User{Login: "ektov", Email: "ektov@mego.travel"},
		Reviewers:  []git2.User{{"admin", "admin@mego.travel"}},
	}

	return res, nil
}

func (p *pullRequestService) GetOpened(repo string) ([]git2.PullRequest, error) {
	//TODO implement me
	res1, _ := p.Get(repo, "1")
	res2, _ := p.Get(repo, "2")

	return []git2.PullRequest{
		res1, res2,
	}, nil
}

func (p *pullRequestService) GetCommits(repo string, number string) ([]git2.Commit, error) {
	//TODO implement me
	return []git2.Commit{
		{"OTAB-123 = test commit"},
		{"OTAB-123 + test commit 2"},
	}, nil
}

func (p *pullRequestService) Create(pr *git2.PullRequest) error {
	//TODO implement me

	createdPrNum := strconv.Itoa(rand.Intn(500))

	pr.Number = createdPrNum
	pr.State = "opened"

	return nil
}
