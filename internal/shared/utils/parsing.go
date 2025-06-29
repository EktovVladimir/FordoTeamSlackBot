package utils

import (
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"regexp"
	"strconv"
	"strings"
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

func ParsePullRequestRefFromUrlTextMany(text string) ([]*models.PullRequestRef, error) {
	words := strings.Fields(text)

	urls := make([]string, 0)
	for _, word := range words {
		if prRegexp.MatchString(word) {
			urls = append(urls, word)
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("no GitHub PR URLs found in text")
	}

	return ParsePullRequestRefFromUrlMany(urls)
}
