package chain

import (
	"chain/github"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ErrLoopedChain = fmt.Errorf("chain has a loop")

type repoService interface {
	GetPullRequest(number uint) (*github.PullRequest, error)
	ListPullRequests() ([]*github.PullRequest, error)
}

type chainHandler struct {
	repoService repoService
}

func NewChainHandler() *chainHandler {
	service := github.NewAdaptor()
	return &chainHandler{
		repoService: service,
	}
}

func (o *chainHandler) GetPullRequests() ([]*Pr, error) {
	gitHubPrs, err := o.repoService.ListPullRequests()

	if err != nil {
		return nil, err
	}

	prs := make([]*Pr, 0, len(gitHubPrs))
	for _, pr := range gitHubPrs {
		pull, err := mapPr(pr)
		if err != nil {
			return nil, err
		}
		prs = append(prs, pull)
	}

	return prs, nil
}

func (o *chainHandler) GetChain(number uint) (map[uint]*Pr, error) {
	gitHubPullRequest, err := o.repoService.GetPullRequest(number)

	if err != nil {
		return nil, err
	}

	pr, err := mapPr(gitHubPullRequest)

	if err != nil {
		return nil, err
	}

	chain := map[uint]*Pr{gitHubPullRequest.Number(): pr}

	for findLink(gitHubPullRequest.Body()) != 0 {
		link, err := o.repoService.GetPullRequest(findLink(gitHubPullRequest.Body()))

		if err != nil {
			return nil, err
		}

		if chain[link.Number()] != nil {
			return nil, ErrLoopedChain
		}

		pr, err = mapPr(link)
		if err != nil {
			return nil, err
		}

		chain[link.Number()] = pr
		gitHubPullRequest = link
	}

	return chain, nil
}

func mapPr(pr *github.PullRequest) (*Pr, error) {
	state, err := mapState(pr.State(), pr.Labels())
	if err != nil {
		return nil, err
	}

	link := InitPr(
		pr.Title(),
		pr.Body(),
		pr.Branch(),
		pr.Number(),
		findLink(pr.Body()),
		state,
	)
	return &link, nil
}

func mapState(state github.State, labels []string) (state, error) {
	switch state {
	case github.StateOpen:
		if isBlocked(labels) {
			return blocked, nil
		}
		return open, nil
	case github.StateClosed:
		return closed, nil
	case github.StateMerged:
		if isReleased(labels) {
			return released, nil
		}
		return merged, nil
	default:
		return 0, fmt.Errorf("unexpected state: %s", state)
	}
}

func isBlocked(labels []string) bool {
	for _, label := range labels {
		if strings.EqualFold(label, "DO NOT MERGE") {
			return true
		}
	}
	return false
}

func isReleased(labels []string) bool {
	for _, label := range labels {
		if strings.EqualFold(label, "RELEASED") {
			return true
		}
	}
	return false
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
