package utils

import (
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"regexp"
	"strconv"
)

var (
	prRegexp = regexp.MustCompile(`github\.com/([^/]+)/([^/]+)/pull/(\d+)`)
)

func ParsePullRequestRefFromUrl(url string) (*models.PullRequestRef, error) {
	matches := prRegexp.FindStringSubmatch(url)
	if len(matches) < 4 {
		return nil, errors.New("invalid GitHub Pull Request URL")
	}

	number, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid PR number: %v", err)
	}

	return models.NewPullRequestRef(matches[1], matches[2], number), nil
}

func ParsePullRequestRefFromUrlMany(urls []string) ([]*models.PullRequestRef, error) {
	res := make([]*models.PullRequestRef, 0)
	for _, url := range urls {
		prRef, err := ParsePullRequestRefFromUrl(url)
		if err != nil {
			return nil, err
		}
		res = append(res, prRef)
	}

	return res, nil
}
