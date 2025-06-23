package chain

import (
	"chain/github"
	"fmt"
)

var (
	ErrFailedToFetch = fmt.Errorf("failed to fetch")
	ErrFailedToMap   = fmt.Errorf("failed to map pull request")
	ErrLoopedChain   = fmt.Errorf("chain has a loop")
)

type gitHubAdaptor interface {
	GetPr(number uint) (*github.PullRequest, error)
	ListOpenPrs() ([]*github.PullRequest, error)
}

type orchestrator struct {
	gitHubAdaptor gitHubAdaptor
}

func InitOrchestrator() *orchestrator {
	adaptor := github.NewAdaptor()
	return &orchestrator{
		gitHubAdaptor: adaptor,
	}
}

func (o *orchestrator) ListOpenPrs() ([]*Pr, error) {
	gitHubPrs, err := o.gitHubAdaptor.ListOpenPrs()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	prs := make([]*Pr, 0, len(gitHubPrs))
	for _, pr := range gitHubPrs {
		mapped, err := mapPr(pr)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrFailedToMap, err)
		}
		prs = append(prs, mapped)
	}

	return prs, nil
}

func (o *orchestrator) GetChain(number uint) (map[uint]*Pr, error) {
	gitHubPr, err := o.gitHubAdaptor.GetPr(number)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	pr, err := mapPr(gitHubPr)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToMap, err)
	}

	chain := map[uint]*Pr{gitHubPr.Number(): pr}

	for findLink(gitHubPr.Body()) != 0 {
		link, err := o.gitHubAdaptor.GetPr(findLink(gitHubPr.Body()))

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
		gitHubPr = link
	}

	return chain, nil
}
