package slack

type Thread struct {
	Ts          string
	IsBot       bool
	Text        string
	Attachments []Attachment
	Blocks      []Block
}
