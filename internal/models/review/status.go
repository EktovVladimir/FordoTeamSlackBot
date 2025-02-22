package review

type Status struct {
	Value    string
	approved bool
	done     bool
}

func (r *Status) Approved() bool {
	return r.approved
}

func (r *Status) Approve(value bool) {
	r.approved = value
}

func (r *Status) IsDone() bool {
	return r.done
}

func (r *Status) SetDone(value bool) {
	r.done = value
}
