package chain

import (
	"chain/github"
	"fmt"
	"slices"
)

type gitHubAdaptor interface {
	GetPr(number uint) (*github.PullRequest, error)
	ListOpenPrs() ([]*github.PullRequest, error)
}

type Orchestrator struct {
	prs           map[uint]*Pr
	gitHubAdaptor gitHubAdaptor
	targetLabel   string
}

func InitOrchestrator(label string) *Orchestrator {
	adaptor := github.NewAdaptor()
	return &Orchestrator{
		prs:           make(map[uint]*Pr),
		gitHubAdaptor: adaptor,
		targetLabel:   label,
	}
}

func (o *Orchestrator) ListOpenPrs() ([]*Pr, error) {
	gitHubPrs, err := o.gitHubAdaptor.ListOpenPrs()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	prs := make([]*Pr, 0, len(gitHubPrs))
	for _, pr := range gitHubPrs {
		linkId := findLinkId(pr.Body())
		link, err := o.getLink(linkId)
		if err != nil {
			return nil, err
		}
		mapped, err := mapGitHubPullRequest(pr, link)
		if err != nil {
			return nil, err
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
		number = pr.LinkId()
	}

	return linkedPrs, nil
}

func (o *Orchestrator) getPr(number uint) (*Pr, error) {
	if pr := o.prs[number]; pr != nil {
		return pr, nil
	}

	gitHubPr, err := o.gitHubAdaptor.GetPr(number)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	linkId := findLinkId(gitHubPr.Body())
	link, err := o.getLink(linkId)
	if err != nil {
		return nil, err
	}

	mapped, err := mapGitHubPullRequest(gitHubPr, link)
	if err != nil {
		return nil, err
	}

	o.prs[mapped.Id()] = mapped
	return mapped, nil
}

func (o *Orchestrator) getLink(linkId uint) (*Link, error) {
	if linkId == 0 {
		return nil, nil
	}

	if o.prs[linkId] != nil {
		return &Link{
			id:             linkId,
			hasTargetLabel: o.prs[linkId].HasLabel(o.targetLabel),
		}, nil
	}

	link, err := o.gitHubAdaptor.GetPr(linkId)
	if err != nil {
		return nil, fmt.Errorf("%w %d: %w", ErrFailedToFetch, linkId, err)
	}

	hasTargetLabel := false
	if link.Labels() != nil {
		if slices.Contains(link.Labels(), o.targetLabel) {
			hasTargetLabel = true
		}
	}

	return &Link{
		id:             linkId,
		hasTargetLabel: hasTargetLabel,
	}, nil
}
