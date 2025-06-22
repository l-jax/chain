package tui

type Link struct {
	title       string
	description string
	id          uint
	linkid      uint
}

func NewLink(title, description string, id, linkid uint) Link {
	return Link{
		title:       title,
		description: description,
		id:          id,
		linkid:      linkid,
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
