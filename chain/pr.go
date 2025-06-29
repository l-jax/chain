package chain

import "strings"

type Pr struct {
	title  string
	body   string
	branch string
	labels []string
	state  state
	id     uint
	link   *Link
}

type Link struct {
	id             uint
	hasTargetLabel bool
}

func NewPr(title, body, branch string, labels []string, id uint, state state, link *Link) *Pr {
	return &Pr{
		title:  title,
		body:   body,
		branch: branch,
		labels: labels,
		id:     id,
		state:  state,
		link:   link,
	}
}

func (p Pr) Title() string  { return p.title }
func (p Pr) Branch() string { return p.branch }
func (p Pr) State() state   { return p.state }
func (p Pr) Id() uint       { return p.id }
func (p Pr) Body() string   { return p.body }

func (p Pr) Labels() []string {
	if p.labels == nil {
		return []string{}
	}
	return p.labels
}

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

func (p Pr) HasLabel(label string) bool {
	if p.labels == nil {
		return false
	}
	for _, l := range p.labels {
		if strings.EqualFold(l, label) {
			return true
		}
	}
	return false
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
