package chain

import (
	"chain/github"
	"errors"
)

func getPullRequests() ([]Pull, error) {
	prs, err := github.ListActivePrs()
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
