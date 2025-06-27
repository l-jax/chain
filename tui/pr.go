package tui

type pr struct {
	id          uint
	title       string
	description string
	text        string
	label       string
	dependsOn   uint
}

func newPr(id uint, title string, description string, text string, label string, dependsOn uint) *pr {
	return &pr{
		id:          id,
		title:       title,
		description: description,
		text:        text,
		label:       label,
		dependsOn:   dependsOn,
	}
}

func (i pr) FilterValue() string { return i.title }
func (i pr) Title() string       { return i.title }
func (i pr) Description() string { return i.description }
func (i pr) Text() string        { return i.text }
func (i pr) Label() string       { return i.label }
func (i pr) Id() uint            { return i.id }
func (i pr) DependsOn() uint     { return i.dependsOn }
