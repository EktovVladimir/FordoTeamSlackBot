package gh_service

import (
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/google/go-github/v72/github"
	"golang.org/x/net/context"
	"maps"
	"slices"
)

var (
	ErrReviewersNotFound = errors.New("reviewers not found")
)

type PullRequestService service

func NewPullRequestService(ghClient githubPullRequestService) *PullRequestService {
	return &PullRequestService{ghClient}
}

func (s *PullRequestService) GetReviewers(ctx context.Context, prRef *models.PullRequestRef) ([]*models.GhReviewer, error) {
	reviewers, _, err := s.ghClient.ListReviewers(ctx, prRef.Owner, prRef.Repo, prRef.Number, &github.ListOptions{
		PerPage: 10,
		Page:    1,
	})
	if err != nil {
		return nil, err
	}

	res := make([]*models.GhReviewer, 0)

	if reviewers.Users == nil {
		return res, ErrReviewersNotFound
	}

	for _, r := range reviewers.Users {
		res = append(res, &models.GhReviewer{
			Email: r.GetEmail(),
			Login: r.GetLogin(),
		})
	}

	return res, nil
}

func (s *PullRequestService) GetReviewersMany(ctx context.Context, prRefs []*models.PullRequestRef) ([]*models.GhReviewer, error) {

	revMap := make(map[string]*models.GhReviewer)
	for _, prRef := range prRefs {
		ghReviews, err := s.GetReviewers(ctx, prRef)
		if err != nil {
			if errors.Is(err, ErrReviewersNotFound) {
				continue
			} else {
				return nil, err
			}
		}

		for _, ghReview := range ghReviews {
			revMap[ghReview.Email] = ghReview
		}
	}

	res := make([]*models.GhReviewer, 0)
	if len(revMap) != 0 {
		res = slices.Collect(maps.Values(revMap))
	}

	return res, nil
}
