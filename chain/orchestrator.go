package chain

import (
	"chain/github"
	"fmt"
)

type gitHubAdaptor interface {
	GetPr(number uint) (*github.PullRequest, error)
	ListOpenPrs() ([]*github.PullRequest, error)
}

type Orchestrator struct {
	gitHubAdaptor gitHubAdaptor
}

func InitOrchestrator() *Orchestrator {
	adaptor := github.NewAdaptor()
	return &Orchestrator{
		gitHubAdaptor: adaptor,
	}
}

func (o *Orchestrator) ListOpenPrs() ([]*Pr, error) {
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

func (o *Orchestrator) GetPrsLinkedTo(number uint) (map[uint]*Pr, error) {
	linkedPrs := make(map[uint]*Pr)

	for number != 0 {
		pr, err := o.getPr(number)

		if err != nil {
			return nil, err
		}

		if linkedPrs[pr.Id()] != nil {
			return nil, ErrLoopedChain
		}

		linkedPrs[pr.Id()] = pr

		number = findLinkedPr(pr.Body())
	}

	return linkedPrs, nil
}

func (o *Orchestrator) getPr(number uint) (*Pr, error) {
	gitHubPr, err := o.gitHubAdaptor.GetPr(number)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	pr, err := mapPr(gitHubPr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToMap, err)
	}
	return pr, nil
}
