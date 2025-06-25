package review_handler

type RequestReviewRequest struct {
	Email     string   `json:"email" validate:"required,email"`
	ChannelId string   `json:"channel_id" validate:"required,min=1"`
	PrUrls    []string `json:"pr_urls" validate:"required,min=1"`
}
