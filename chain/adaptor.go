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
	return mapToPull(pr)
}

func (a *ghAdaptor) listPullRequests() ([]*Pull, error) {
	prs, err := a.port.ListActivePrs()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all: %w", err)
	}

	pulls := make([]*Pull, 0, len(prs))
	for _, pr := range prs {
		pull, err := mapToPull(pr)
		if err != nil {
			log.Println(err)
			continue
		}
		pulls = append(pulls, pull)
	}
	return pulls, nil
}

func findLink(body string) uint {
	re := regexp.MustCompile(`do not merge until #(\d+)`)
	match := re.FindStringSubmatch(body)

	if match == nil {
		return 0
	}

	link, err := strconv.ParseUint(match[1], 10, 32)
	if err != nil {
		return 0
	}
	return uint(link)
}

func mapToPull(pr *github.GhPullRequest) (*Pull, error) {
	state, err := mapState(pr.State, pr.Labels)

	if err != nil {
		return nil, fmt.Errorf("failed to map pull request %s: %w", pr.HeadRefName, err)
	}

	link := findLink(pr.Body)

	return NewPull(pr.Title, pr.HeadRefName, pr.Body, state, pr.Number, link), nil
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
