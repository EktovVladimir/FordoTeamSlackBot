package github

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
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

func (p *pullRequestService) Get(repo string, number string) (git.PullRequest, error) {
	//TODO implement me
	res := git.PullRequest{
		Number:     number,
		Repo:       repo,
		Branch:     "develop",
		BaseBranch: "master",
		State:      "opened",
		Requester:  git.User{Login: "ektov", Email: "ektov@mego.travel"},
		Reviewers:  []git.User{{"admin", "admin@mego.travel"}},
	}

	return res, nil
}

func (p *pullRequestService) GetOpened(repo string) ([]git.PullRequest, error) {
	//TODO implement me
	res1, _ := p.Get(repo, "1")
	res2, _ := p.Get(repo, "2")

	return []git.PullRequest{
		res1, res2,
	}, nil
}

func (p *pullRequestService) GetCommits(repo string, number string) ([]git.Commit, error) {
	//TODO implement me
	return []git.Commit{
		{"OTAB-123 = test commit"},
		{"OTAB-123 + test commit 2"},
	}, nil
}

func (p *pullRequestService) Create(pr *git.PullRequest) error {
	//TODO implement me

	createdPrNum := strconv.Itoa(rand.Intn(500))

	pr.Number = createdPrNum
	pr.State = "opened"

	return nil
}
