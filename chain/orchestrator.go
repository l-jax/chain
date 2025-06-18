package chain

import (
	"chain/github"
	"fmt"
)

var ErrLoopedChain = fmt.Errorf("chain has a loop")

type orchestrator struct {
	adaptor adaptor
}

func NewOrchestrator() *orchestrator {
	adaptor := &ghAdaptor{
		port: github.GhPort{},
	}
	return &orchestrator{
		adaptor: adaptor,
	}
}

func (o *orchestrator) GetPullRequests() ([]*Pull, error) {
	pulls, err := o.adaptor.listPullRequests()

	if err != nil {
		return nil, err
	}

	return pulls, nil
}

func (o *orchestrator) GetChain(number uint) (map[uint]*Pull, error) {
	pull, err := o.adaptor.getPullRequest(number)

	if err != nil {
		return nil, err
	}

	chain := map[uint]*Pull{pull.Number(): pull}

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
