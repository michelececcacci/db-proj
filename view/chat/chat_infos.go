package chat

type chatInfo struct {
	id          int32
	name        string
	description string
}

func (c chatInfo) Title() string {
	return c.name
}

func (c chatInfo) FilterValue() string {
	return c.name
}

func (c chatInfo) Description() string {
	return c.description
}
