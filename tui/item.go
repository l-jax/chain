package tui

type Item struct {
	id             uint
	title          string
	description    string
	text           string
	state          string
	dependsOn      uint
	blocked        bool
	hasTargetLabel bool
}

func newItem(
	id uint,
	title string,
	description string,
	text string,
	state string,
	dependsOn uint,
	blocked bool,
	hasTargetLabel bool,
) *Item {
	return &Item{
		id:             id,
		title:          title,
		description:    description,
		text:           text,
		state:          state,
		dependsOn:      dependsOn,
		blocked:        blocked,
		hasTargetLabel: hasTargetLabel,
	}
}

func (i Item) FilterValue() string  { return i.title }
func (i Item) Title() string        { return i.title }
func (i Item) Description() string  { return i.description }
func (i Item) Text() string         { return i.text }
func (i Item) State() string        { return i.state }
func (i Item) Id() uint             { return i.id }
func (i Item) DependsOn() uint      { return i.dependsOn }
func (i Item) Blocked() bool        { return i.blocked }
func (i Item) HasTargetLabel() bool { return i.hasTargetLabel }
