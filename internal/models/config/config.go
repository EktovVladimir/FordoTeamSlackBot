package config

type ConfigItem struct {
	key            string
	channelPattern string
	Value          string
}

func (c ConfigItem) Key() string {
	return c.key
}

func (c ConfigItem) ChannelPattern() string {
	return c.channelPattern
}

func NewConfigItem(key string, channel string, value string) *ConfigItem {
	return &ConfigItem{
		key:            key,
		channelPattern: channel,
		Value:          value,
	}
}
