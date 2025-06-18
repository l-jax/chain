package chain

type PullRequest struct {
	title  string
	branch string
	body   string
	state  State
	number uint
	chain  uint
}

func NewPullRequest(title, branch, body string, state State, number uint, chain uint) *PullRequest {
	return &PullRequest{state: state, title: title, body: body, branch: branch, number: number, chain: chain}
}

func (p PullRequest) Title() string  { return p.title }
func (p PullRequest) Branch() string { return p.branch }
func (p PullRequest) Body() string   { return p.body }
func (p PullRequest) State() State   { return p.state }
func (p PullRequest) Number() uint   { return p.number }
func (p PullRequest) Chain() uint    { return p.chain }

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
