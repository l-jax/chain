package chain

import (
	"chain/github"
	"fmt"
	"log"
)

type adaptor interface {
	getPullRequest() (*Pull, error)
	listPullRequests() ([]*Pull, error)
}

type ghAdaptor struct {
	port github.Port
}

func (a *ghAdaptor) getPullRequest(branch string) (*Pull, error) {
	pr, err := a.port.GetPr(branch)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch branch %s: %w", branch, err)
	}
	return mapPr(pr)
}

func (a *ghAdaptor) listPullRequests() ([]*Pull, error) {
	prs, err := a.port.ListActivePrs()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all: %w", err)
	}

	var pullRequests []*Pull
	for _, pr := range prs {
		pull, err := mapPr(pr)
		if err != nil {
			log.Println(err)
			continue
		}
		pullRequests = append(pullRequests, pull)
	}
	return pullRequests, nil
}

func mapPr(pr *github.GhPullRequest) (*Pull, error) {
	state, err := mapState(pr.State, pr.Labels)

	if err != nil {
		return nil, fmt.Errorf("failed to map pull request %s: %w", pr.HeadRefName, err)
	}

	return NewPull(pr.Title, pr.HeadRefName, state, nil), nil
}

func mapState(state string, labels []github.GhLabel) (State, error) {
	if isReleased(labels) {
		return StateReleased, nil
	}

	switch state {
	case "OPEN":
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		return StateMerged, nil
	default:
		return 0, fmt.Errorf("unexpected state: %s", state)
	}
}

func isReleased(labels []github.GhLabel) bool {
	for _, label := range labels {
		if label.Name == "RELEASED" {
			return true
		}
	}
	return false
}
