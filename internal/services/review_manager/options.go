package review_manager

type RequestReviewOption func(*requestReviewConfig) error

type requestReviewConfig struct {
	AsUser             bool
	preferredReviewers []string
	jiraProject        string
}

func defaultCreateThreadConfig() *requestReviewConfig {
	return &requestReviewConfig{
		AsUser:             false,
		preferredReviewers: []string{},
		jiraProject:        "OTA[A-Z]+",
	}
}

func (cfg *requestReviewConfig) setOptions(opts ...RequestReviewOption) error {
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return err
		}
	}
	return nil
}

func WithPreferredReviewers(preferredReviewers []string) RequestReviewOption {
	return func(config *requestReviewConfig) error {
		config.preferredReviewers = preferredReviewers
		return nil
	}
}

func WithJiraProject(jiraProject string) RequestReviewOption {
	return func(config *requestReviewConfig) error {
		config.jiraProject = jiraProject
		return nil
	}
}

func WithAsUser() RequestReviewOption {
	return func(config *requestReviewConfig) error {
		config.AsUser = true
		return nil
	}
}
