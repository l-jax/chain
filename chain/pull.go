package chain

type Pull struct {
	title  string
	branch string
	state  State
	chain  *Pull
}

func NewPull(title, branch string, state State, chain *Pull) Pull {
	return Pull{state: state, title: title, branch: branch, chain: chain}
}

func (p Pull) Title() string  { return p.title }
func (p Pull) Branch() string { return p.branch }
func (p Pull) State() State   { return p.state }

type State uint

const (
	StateOpen State = iota
	StateMerged
	StateClosed
	StateReleased
)

var stateName = map[State]string{
	StateOpen:     "open",
	StateMerged:   "merged",
	StateClosed:   "closed",
	StateReleased: "released",
}

func (s State) String() string {
	return stateName[s]
}
