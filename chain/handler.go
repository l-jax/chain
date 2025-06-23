package chain

import (
	"chain/github"
	"fmt"
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

func (o *chainHandler) GetPullRequests() ([]*github.PullRequest, error) {
	pulls, err := o.repoService.ListPullRequests()

	if err != nil {
		return nil, err
	}

	return pulls, nil
}

func (o *chainHandler) GetChain(number uint) (map[uint]*github.PullRequest, error) {
	pull, err := o.repoService.GetPullRequest(number)

	if err != nil {
		return nil, err
	}

	chain := map[uint]*github.PullRequest{pull.Number(): pull}

	for pull.Chain() != 0 {
		link, err := o.repoService.GetPullRequest(pull.Chain())

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
