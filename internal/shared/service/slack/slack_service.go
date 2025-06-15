package slack

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	messenger2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/messenger"
	"strconv"
	"time"
)

type slackService struct {
	config config.Slack
}

func New(config config.Slack) *slackService {
	//TODO создание клиента GH
	return &slackService{
		config: config,
	}
}

func (service *slackService) GetUser(email string) (messenger2.User, error) {
	//TODO implement me
	return messenger2.User{
		Id:   "default",
		Name: "ektov",
	}, nil
}

func (service *slackService) GetThread(channel string, ts string) (messenger2.Thread, error) {
	//TODO implement me
	res := messenger2.Thread{
		Ts:     ts,
		UserId: "dummy",
		BotId:  "dummy",
		Text:   "example",
		Attachments: []messenger2.Attachment{
			{"attachment text", "default"},
		},
	}
	return res, nil
}

func (service *slackService) SearchThreadLast(channel string, searchCriteria string) (messenger2.Thread, error) {
	//TODO implement me
	return service.GetThread(channel, "1747948358.528789")
}

func (service *slackService) SearchCommentLast(channel string, ts string, searchCriteria string) (messenger2.Comment, error) {
	//TODO implement me
	res := messenger2.Comment{Ts: "1747948358.528789", Text: "example"}
	return res, nil
}

func (service *slackService) CreateThread(channel string, thread *messenger2.Thread) error {
	//TODO implement me
	ts := generateSlackTimestamp(time.Now())
	thread.Ts = ts
	thread.UserId = "dummy"
	thread.BotId = "dummy"

	return nil
}

func (service *slackService) UpdateThread(channel string, thread *messenger2.Thread) error {
	//TODO implement me
	return nil
}

func (service *slackService) CreateComment(channel string, ts string, comment *messenger2.Comment) error {
	//TODO implement me
	commentTs := generateSlackTimestamp(time.Now())
	comment.Ts = commentTs
	return nil
}

func (service *slackService) UpdateComment(channel string, ts string, comment *messenger2.Comment) error {
	//TODO implement me
	return nil
}

func generateSlackTimestamp(t time.Time) string {
	seconds := t.Unix()
	microseconds := t.Nanosecond() / 1000
	return strconv.FormatInt(seconds, 10) + "." + strconv.Itoa(microseconds)
}
