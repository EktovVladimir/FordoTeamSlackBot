package messenger

type Thread struct {
	Ts          string
	UserId      string
	BotId       string
	Text        string
	Attachments []Attachment
	Blocks      []Block
}
