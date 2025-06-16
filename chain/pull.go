package chain

type Pull struct {
	title  string
	branch string
	body   string
	state  State
	chain  *Pull
}

func NewPull(title, branch, body string, state State) *Pull {
	return &Pull{state: state, title: title, body: body, branch: branch}
}

func (p Pull) Title() string  { return p.title }
func (p Pull) Branch() string { return p.branch }
func (p Pull) Body() string   { return p.body }
func (p Pull) State() State   { return p.state }
func (p Pull) Chain() *Pull   { return p.chain }

type State uint

const (
	StateOpen State = iota
	StateBlocked
	StateMerged
	StateClosed
	StateReleased
)

var stateName = map[State]string{
	StateOpen:     "open",
	StateBlocked:  "blocked",
	StateMerged:   "merged",
	StateClosed:   "closed",
	StateReleased: "released",
}

func (s State) String() string {
	return stateName[s]
}
