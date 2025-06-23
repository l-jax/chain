package chain

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

type Pr struct {
	title  string
	body   string
	branch string
	state  state
	id     uint
	linkid uint
}

func InitPr(title, body, branch string, id, linkid uint, state state) Pr {
	return Pr{
		title:  title,
		body:   body,
		id:     id,
		linkid: linkid,
		state:  state,
		branch: branch,
	}
}

func (l Pr) FilterValue() string { return l.title }
func (l Pr) Title() string       { return l.branch }
func (l Pr) Description() string { return l.state.String() }
func (l Pr) Branch() string      { return l.branch }
func (l Pr) State() state        { return l.state }
func (l Pr) Id() uint            { return l.id }
func (l Pr) LinkId() uint        { return l.linkid }
func (l Pr) Body() string        { return l.body }
