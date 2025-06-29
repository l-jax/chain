package chain

type Pr struct {
	title  string
	body   string
	branch string
	state  state
	id     uint
	link   *Link
}

type Link struct {
	id             uint
	hasTargetLabel bool
}

func NewPr(title, body, branch string, id uint, state state, link *Link) *Pr {
	return &Pr{
		title:  title,
		body:   body,
		id:     id,
		state:  state,
		branch: branch,
		link:   link,
	}
}

func (p Pr) Title() string  { return p.title }
func (p Pr) Branch() string { return p.branch }
func (p Pr) State() state   { return p.state }
func (p Pr) Id() uint       { return p.id }
func (p Pr) Body() string   { return p.body }
func (p Pr) LinkId() uint {
	if p.link == nil {
		return 0
	}
	return p.link.id
}
func (p Pr) Blocked() bool {
	if p.link == nil {
		return false
	}
	return !p.link.hasTargetLabel
}

type state uint

const (
	draft state = iota
	open
	merged
	released
	closed
)

func (s state) String() string {
	switch s {
	case state(draft):
		return "draft"
	case state(open):
		return "open"
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
