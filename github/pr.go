package github

type PullRequest struct {
	title  string
	branch string
	body   string
	state  State
	labels []string
	number uint
}

func NewPullRequest(title, branch, body string, state State, labels []string, number uint) *PullRequest {
	return &PullRequest{
		title:  title,
		branch: branch,
		body:   body,
		state:  state,
		labels: labels,
		number: number,
	}
}

func (p PullRequest) Title() string    { return p.title }
func (p PullRequest) Branch() string   { return p.branch }
func (p PullRequest) Body() string     { return p.body }
func (p PullRequest) State() State     { return p.state }
func (p PullRequest) Labels() []string { return p.labels }
func (p PullRequest) Number() uint     { return p.number }

type State uint

const (
	StateDraft State = iota
	StateOpen
	StateMerged
	StateClosed
)

var stateName = map[State]string{
	StateDraft:  "draft",
	StateOpen:   "open",
	StateMerged: "merged",
	StateClosed: "closed",
}

func (s State) String() string {
	return stateName[s]
}
