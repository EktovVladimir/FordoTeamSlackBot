package request_source

type Workflow struct {
	Repo   string
	Branch string
	Team   string
}

func (w Workflow) GetKey() string {
	return w.Team
}
