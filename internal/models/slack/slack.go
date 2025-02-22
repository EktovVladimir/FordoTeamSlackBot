package slack

type TheadInfo struct {
	Code    string
	Channel ChannelInfo
	Author  User
}

type ChannelInfo struct {
	Name string
	Code string
}

type User struct {
	Code string
}
