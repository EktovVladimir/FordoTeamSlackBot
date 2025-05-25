package abstract

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"

type MessengerService interface {
	GetThread(channel string, ts string) (messenger.Thread, error)
	SearchThreadLast(channel string, searchCriteria string) (messenger.Thread, error)
	SearchCommentLast(channel string, ts string, searchCriteria string) (messenger.Comment, error)
	CreateThread(channel string, thread *messenger.Thread) error
	UpdateThread(channel string, thread *messenger.Thread) error
	CreateComment(channel string, ts string, comment *messenger.Comment) error
	UpdateComment(channel string, ts string, comment *messenger.Comment) error
}
