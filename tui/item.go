package tui

type Item struct {
	id          uint
	title       string
	description string
	text        string
	label       string
	dependsOn   uint
}

func NewItem(id uint, title string, description string, text string, label string, dependsOn uint) *Item {
	return &Item{
		id:          id,
		title:       title,
		description: description,
		text:        text,
		label:       label,
		dependsOn:   dependsOn,
	}
}

func (i Item) FilterValue() string { return i.title }
func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) Text() string        { return i.text }
func (i Item) Label() string       { return i.label }
func (i Item) Id() uint            { return i.id }
func (i Item) DependsOn() uint     { return i.dependsOn }
