package chain

import (
	"chain/github"
)

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

func (o *orchestrator) GetChain(number uint) ([]*Pull, error) {
	pull, err := o.adaptor.getPullRequest(number)
	chain := []*Pull{pull}

	if err != nil {
		return nil, err
	}

	for pull.Chain() != 0 {
		link, err := o.adaptor.getPullRequest(pull.Chain())

		if err != nil {
			return nil, err
		}

		chain = append(chain, link)
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
