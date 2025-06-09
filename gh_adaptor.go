package main

import (
	"errors"
	"github.com/charmbracelet/bubbles/list"
)

func getPullRequests() ([]list.Item, error) {
	prs, err := listActivePrs()
	if err != nil {
		return nil, err
	}

	var pullRequests []list.Item
	for _, pr := range prs {
		state, err := mapState(pr.State)
		if err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, NewPull(pr.Title, pr.HeadRefName, state, nil))
	}
	return pullRequests, nil
}

func mapState(state string) (state, error) {
	switch state {
	case "OPEN":
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		return StateMerged, nil
	default:
		return -1, errors.New("unexpected state: " + state)
	}
}
