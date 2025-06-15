package helpers

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
)

func GetPullRequestUrl(pr git.PullRequest) string {
	return fmt.Sprintf("https://github.com/%s/%s/pull/%s", pr.Owner, pr.Repo, pr.Number)
}
