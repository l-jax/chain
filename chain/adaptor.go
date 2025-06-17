package chain

import (
	"chain/github"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type adaptor interface {
	getPullRequest(number uint) (*Pull, error)
	listPullRequests() ([]*Pull, error)
}

type ghAdaptor struct {
	port github.Port
}

func (a *ghAdaptor) getPullRequest(number uint) (*Pull, error) {
	pr, err := a.port.GetPr(fmt.Sprintf("%d", number))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pull %d: %w", number, err)
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

func findLinkedPrNumberInBody(body string) uint {
	re := regexp.MustCompile(`do not merge until #(\d+)`)
	match := re.FindStringSubmatch(body)

	if match == nil {
		return 0
	}

	number, err := strconv.ParseUint(match[1], 10, 32)
	if err != nil {
		return 0
	}
	return uint(number)
}

func mapPr(pr *github.GhPullRequest) (*Pull, error) {
	state, err := mapState(pr.State, pr.Labels)

	if err != nil {
		return nil, fmt.Errorf("failed to map pull request %s: %w", pr.HeadRefName, err)
	}

	linkedPrNumber := findLinkedPrNumberInBody(pr.Body)

	return NewPull(pr.Title, pr.HeadRefName, pr.Body, state, pr.Number, linkedPrNumber), nil
}

func mapState(state string, labels []github.GhLabel) (State, error) {
	switch state {
	case "OPEN":
		if isBlocked(labels) {
			return StateBlocked, nil
		}
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		if isReleased(labels) {
			return StateReleased, nil
		}
		return StateMerged, nil
	default:
		return 0, fmt.Errorf("unexpected state: %s", state)
	}
}

func isBlocked(labels []github.GhLabel) bool {
	for _, label := range labels {
		if label.Name == "DO NOT MERGE" {
			return true
		}
	}
	return false
}

func isReleased(labels []github.GhLabel) bool {
	for _, label := range labels {
		if label.Name == "RELEASED" {
			return true
		}
	}
	return false
}
