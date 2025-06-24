package chain

type Pr struct {
	title  string
	body   string
	branch string
	state  state
	id     uint
	linkid uint
}

func NewPr(title, body, branch string, id, linkid uint, state state) *Pr {
	return &Pr{
		title:  title,
		body:   body,
		id:     id,
		linkid: linkid,
		state:  state,
		branch: branch,
	}
}

func (p Pr) Title() string  { return p.title }
func (p Pr) Branch() string { return p.branch }
func (p Pr) State() state   { return p.state }
func (p Pr) Id() uint       { return p.id }
func (p Pr) LinkId() uint   { return p.linkid }
func (p Pr) Body() string   { return p.body }

type state uint

const (
	blocked state = iota
	open
	merged
	released
	closed
)

func (s state) String() string {
	switch s {
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
