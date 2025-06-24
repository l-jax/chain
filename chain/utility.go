package chain

import (
	"chain/github"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	releasedLabel   = "RELEASED"
	blockedLabel    = "DO NOT MERGE"
	linkedPrPattern = `do not merge until #(\d+)`
)

func mapPr(pr *github.PullRequest) (*Pr, error) {
	state, err := mapState(pr.State(), pr.Labels())
	if err != nil {
		return nil, err
	}

	mapped := NewPr(
		pr.Title(),
		pr.Body(),
		pr.Branch(),
		pr.Number(),
		findLinkedPr(pr.Body()),
		state,
	)
	return mapped, nil
}

func mapState(state github.State, labels []string) (state, error) {
	switch state {
	case github.StateOpen:
		if labelsContains(labels, blockedLabel) {
			return blocked, nil
		}
		return open, nil
	case github.StateClosed:
		return closed, nil
	case github.StateMerged:
		if labelsContains(labels, releasedLabel) {
			return released, nil
		}
		return merged, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnexpectedState, state)
	}
}

func labelsContains(labels []string, label string) bool {
	for _, l := range labels {
		if strings.EqualFold(l, label) {
			return true
		}
	}
	return false
}

func findLinkedPr(body string) uint {
	re := regexp.MustCompile(linkedPrPattern)
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
