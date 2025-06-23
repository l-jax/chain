package tui

type state uint

const (
	blocked state = iota
	open
	merged
	released
	closed
)

func (l state) String() string {
	switch l {
	case state(open):
		return "open"
	case state(blocked):
		return "blocked"
	case state(merged):
		return "merged"
	case state(released):
		return "released"
	case state(closed):
		return "closed"
	default:
		return "unknown"
	}
}

type pr struct {
	title  string
	body   string
	branch string
	state  state
	id     uint
	linkid uint
}

func InitPr(title, body, branch string, id, linkid uint, state state) pr {
	return pr{
		title:  title,
		body:   body,
		id:     id,
		linkid: linkid,
		state:  state,
		branch: branch,
	}
}

func (l pr) FilterValue() string {
	return l.title
}

func (l pr) Title() string {
	return l.branch
}

func (l pr) Description() string {
	return l.state.String()
}

func (l pr) Branch() string {
	return l.branch
}

func (l pr) State() state {
	return l.state
}

func (l pr) Id() uint {
	return l.id
}

func (l pr) LinkId() uint {
	return l.linkid
}

func (l pr) Body() string {
	return l.body
}
