package tui

type label uint

const (
	open label = iota
	blocked
	merged
	released
	closed
)

func (l label) String() string {
	switch l {
	case label(open):
		return "open"
	case label(blocked):
		return "blocked"
	case label(merged):
		return "merged"
	case label(released):
		return "released"
	case label(closed):
		return "closed"
	default:
		return "unknown"
	}
}

type Link struct {
	title  string
	body   string
	branch string
	label  label
	id     uint
	linkid uint
}

func NewLink(title, body, branch string, id, linkid uint, label label) Link {
	return Link{
		title:  title,
		body:   body,
		id:     id,
		linkid: linkid,
	}
}

func (l Link) FilterValue() string {
	return l.title
}

func (l Link) Title() string {
	return l.title
}

func (l Link) Description() string {
	return l.label.String()
}

func (l Link) Branch() string {
	return l.branch
}

func (l Link) Label() label {
	return l.label
}

func (l Link) Id() uint {
	return l.id
}
