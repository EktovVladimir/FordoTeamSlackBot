package gh_service

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/google/go-github/v72/github"
	"golang.org/x/net/context"
	"regexp"
	"strings"
)

type CommitRetriever service

func NewCommitRetriever(ghClient githubPullRequestService) *CommitRetriever {
	return &CommitRetriever{ghClient}
}

func (c *CommitRetriever) GetAllCommits(ctx context.Context, prRef *models.PullRequestRef) ([]*models.CommitInfo, error) {
	res := make([]*models.CommitInfo, 0)
	page := 1
	for {
		paginateOpt := &github.ListOptions{
			Page:    page,
			PerPage: pageSize,
		}

		commits, _, err := c.ghClient.ListCommits(ctx, prRef.Owner, prRef.Repo, prRef.Number, paginateOpt)
		if err != nil {
			return nil, err
		}

		if len(commits) == 0 {
			break
		}

		for _, commit := range commits {
			if commit.Commit == nil {
				continue
			}

			res = append(res, &models.CommitInfo{
				Message: commit.Commit.GetMessage(),
			})
		}

		page++
	}

	return res, nil
}

func (c *CommitRetriever) GetIssueNumbersMany(ctx context.Context, prRefs []*models.PullRequestRef, issuePattern string) ([]string, error) {
	//Добавим ignore-case, если не был передан
	if !strings.HasPrefix(issuePattern, "(?i)") {
		issuePattern = "(?i)" + issuePattern
	}

	rxp, err := regexp.Compile(issuePattern)
	if err != nil {
		return nil, err
	}

	//issueSet вспомогателен и нужен для проверки уникальности
	//res - результирующий слайс, нужен для сохранения сортировки по дате добавления коммита
	issueSet := make(map[string]struct{})
	res := make([]string, 0)

	for _, prRef := range prRefs {
		commits, err := c.GetAllCommits(ctx, prRef)
		if err != nil {
			return nil, err
		}

		for _, commit := range commits {
			match := rxp.FindString(commit.Message)
			if match == "" {
				continue
			}

			match = strings.Trim(match, " \n\t ")
			match = strings.ToUpper(match)

			_, ok := issueSet[match]
			if !ok {

				issueSet[match] = struct{}{}
				res = append(res, match)
			}
		}
	}

	return res, nil
}

func (c *CommitRetriever) GetIssueNumbers(ctx context.Context, prRef *models.PullRequestRef, issuePattern string) ([]string, error) {
	return c.GetIssueNumbersMany(ctx, []*models.PullRequestRef{prRef}, issuePattern)
}
