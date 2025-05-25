package slack

type Attachment struct {
	Color string
	Text  string
}

func NewSuccessAttachment(text string) Attachment {
	return Attachment{
		Color: "#2EB67D",
		Text:  text,
	}
}

func NewErrorAttachment(text string) Attachment {
	return Attachment{
		Color: "#FF0000",
		Text:  text,
	}
}

func NewInProcessAttachment(text string) Attachment {
	return Attachment{
		Color: "#48C774",
		Text:  text,
	}
}

func NewAttachment(text string) Attachment {
	return Attachment{
		Color: "#DDDDDD",
		Text:  text,
	}
}
