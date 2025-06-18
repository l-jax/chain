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

func GetOpen() []Pull {
	return []Pull{
		{"my pull request", "some-branch", "some body", StateOpen, 1, 0},
		{"code", "some-other-branch", "some body", StateOpen, 2, 0},
	}
}

func GetChain(branch string) []Pull {
	return []Pull{
		{"remove something", "some-branch-123", "some body", StateReleased, 3, 0},
		{"add something", "my-branch", "some body", StateMerged, 2, 3},
		{"do something", branch, "some body", StateOpen, 1, 2},
	}
}
