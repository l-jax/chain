package tui

type Link struct {
	title       string
	description string
}

func NewLink(title, description string) Link {
	return Link{
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
