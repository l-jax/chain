package chain

import (
	"chain/github"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const linkedPrPattern = `(?i)do not merge until #(\d+)`

func mapGitHubPullRequest(pr *github.PullRequest, link *Link) (*Pr, error) {
	state, err := mapState(pr.State())
	if err != nil {
		return nil, err
	}

	mapped := NewPr(
		pr.Title(),
		pr.Body(),
		pr.Branch(),
		pr.Labels(),
		pr.Number(),
		state,
		link,
	)
	return mapped, nil
}

func mapState(state github.State) (state, error) {
	switch state {
	case github.StateDraft:
		return draft, nil
	case github.StateOpen:
		return open, nil
	case github.StateClosed:
		return closed, nil
	case github.StateMerged:
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

func findLinkId(body string) uint {
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
