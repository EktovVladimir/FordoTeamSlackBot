package config

type Group struct {
	name  string
	users []string
}

func (g Group) Users() []string {
	return g.users
}

func (g Group) Name() string {
	return g.name
}

func NewGroup(name string, users []string) *Group {
	return &Group{
		name:  name,
		users: users}
}
