package main

type State int

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

type item struct {
	title   string
	branch  string
	state   State
	chained bool
}

func (i item) Title() string       { return i.branch }
func (i item) Description() string { return i.title }
func (i item) State() State        { return i.state }
func (i item) Chained() bool       { return i.chained }
func (i item) FilterValue() string { return i.title }
