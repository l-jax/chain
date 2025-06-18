package chain

import (
	"chain/github"
	"fmt"
)

var ErrLoopedChain = fmt.Errorf("chain has a loop")

type adaptor interface {
	getPullRequest(number uint) (*PullRequest, error)
	listPullRequests() ([]*PullRequest, error)
}

type orchestrator struct {
	adaptor adaptor
}

func NewOrchestrator() *orchestrator {
	adaptor := &gitHubAdaptor{
		port: github.GitHubPort{},
	}
	return &orchestrator{
		adaptor: adaptor,
	}
}

func (o *orchestrator) GetPullRequests() ([]*PullRequest, error) {
	pulls, err := o.adaptor.listPullRequests()

	if err != nil {
		return nil, err
	}

	return pulls, nil
}

func (o *orchestrator) GetChain(number uint) (map[uint]*PullRequest, error) {
	pull, err := o.adaptor.getPullRequest(number)

	if err != nil {
		return nil, err
	}

	chain := map[uint]*PullRequest{pull.Number(): pull}

	for pull.Chain() != 0 {
		link, err := o.adaptor.getPullRequest(pull.Chain())

		if err != nil {
			return nil, err
		}

		if chain[link.Number()] != nil {
			return nil, ErrLoopedChain
		}

		chain[link.Number()] = link
		pull = link
	}

	return chain, nil
}
