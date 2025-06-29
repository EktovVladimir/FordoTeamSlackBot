package slack_handler

type slackSource struct {
	userId    string
	email     string
	channelId string
}

func (s slackSource) GetChannelId() string {
	return s.channelId
}

func (s slackSource) GetEmail() string {
	return s.email
}

func (s slackSource) GetUserId() string {
	return s.userId
}
