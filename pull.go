package main

type Pull struct {
	title  string
	branch string
	state  state
	chain  *Pull
}

func NewPull(title, branch string, state state, chain *Pull) Pull {
	return Pull{state: state, title: title, branch: branch, chain: chain}
}

func (p Pull) Title() string       { return p.title }
func (p Pull) Description() string { return p.branch }
func (p Pull) FilterValue() string { return p.title }

type state int

const (
	StateOpen state = iota
	StateMerged
	StateClosed
	StateReleased
)

var stateName = map[state]string{
	StateOpen:     "open",
	StateMerged:   "merged",
	StateClosed:   "closed",
	StateReleased: "released",
}

func (s state) String() string {
	return stateName[s]
}
