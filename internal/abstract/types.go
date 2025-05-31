package abstract

type Source interface {
	GetKey() string
}

type Event interface {
	GetInitiator() string
	Print() string
}
