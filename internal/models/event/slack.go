package event

import "fmt"

type SlackCommandEvent struct {
	Command  string
	Args     string
	Channel  string
	UserName string
}

func (e SlackCommandEvent) GetInitiator() string {
	return e.UserName
}

func (e SlackCommandEvent) Print() string {
	return fmt.Sprintf("slack command event: %s %s in %s from %s",
		e.Command, e.Args, e.Channel, e.UserName)
}
