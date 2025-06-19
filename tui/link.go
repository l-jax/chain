package tui

type index uint

const (
	active index = iota
	chain
)

type Link struct {
	index       index
	title       string
	description string
}

func NewLink(title, description string, index index) Link {
	return Link{
		index:       index,
		title:       title,
		description: description,
	}
}

func (l Link) FilterValue() string {
	return l.title
}

func (l Link) Title() string {
	return l.title
}

func (l Link) Description() string {
	return l.description
}
