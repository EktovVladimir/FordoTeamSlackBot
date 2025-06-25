package review_handler

type apiSource struct {
	email     string
	channelId string
}

func (s apiSource) GetChannelId() string {
	return s.channelId
}

func (s apiSource) GetEmail() string {
	return s.email
}
