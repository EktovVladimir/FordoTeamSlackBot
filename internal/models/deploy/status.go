package deploy

type Status struct {
	inProgress bool
	done       bool
	paused     bool
}

func (s *Status) InProgress() bool {
	return s.inProgress
}

func (s *Status) SetInProgress(inProgress bool) {
	s.inProgress = inProgress
}

func (s *Status) Done() bool {
	return s.done
}

func (s *Status) SetDone(done bool) {
	s.done = done
}

func (s *Status) Paused() bool {
	return s.paused
}

func (s *Status) SetPaused(paused bool) {
	s.paused = paused
}
