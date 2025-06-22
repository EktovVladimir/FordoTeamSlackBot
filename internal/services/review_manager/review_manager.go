package review_manager

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"

type ReviewManager struct {
	userRepo repository.UserRepository
}

func (r *ReviewManager) CreateThreadFromPr(email string, prNumber string, options *CodeReviewOptions) error {
	return nil
}

type CodeReviewOptions struct {
	PreferredReviewers []string
}
