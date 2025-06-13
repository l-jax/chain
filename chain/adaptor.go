package chain

import (
	"chain/github"
	"errors"
)

type adaptor interface {
	getPullRequests() ([]Pull, error)
}

type ghAdaptor struct {
	port github.Port
}

func (a *ghAdaptor) getPullRequests() ([]Pull, error) {
	prs, err := a.port.ListPrs("all")
	if err != nil {
		return nil, err
	}

	var pullRequests []Pull
	for _, pr := range prs {
		state, err := mapState(pr.State)
		if err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, NewPull(pr.Title, pr.HeadRefName, state, nil))
	}
	return pullRequests, nil
}

func mapState(state string) (State, error) {
	switch state {
	case "OPEN":
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		return StateMerged, nil
	default:
		return 0, errors.New("unexpected state: " + state)
	}
}
