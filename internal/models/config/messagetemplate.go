package config

type MessageTemplate struct {
	key      string
	Template string
}

func (m MessageTemplate) Key() string {
	return m.key
}

func NewMessageTemplate(key string, template string) *MessageTemplate {
	return &MessageTemplate{key: key, Template: template}
}
