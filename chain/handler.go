package chain

import (
	"chain/github"
	"fmt"
)

var ErrLoopedChain = fmt.Errorf("chain has a loop")

type repoService interface {
	getPullRequest(number uint) (*PullRequest, error)
	listPullRequests() ([]*PullRequest, error)
}

type chainHandler struct {
	repoService repoService
}

func NewChainHandler() *chainHandler {
	service := &gitHubService{
		gitHubClient: github.GitHubPort{},
	}
	return &chainHandler{
		repoService: service,
	}
}

func (o *chainHandler) GetPullRequests() ([]*PullRequest, error) {
	pulls, err := o.repoService.listPullRequests()

	if err != nil {
		return nil, err
	}

	return pulls, nil
}

func (o *chainHandler) GetChain(number uint) (map[uint]*PullRequest, error) {
	pull, err := o.repoService.getPullRequest(number)

	if err != nil {
		return nil, err
	}

	chain := map[uint]*PullRequest{pull.Number(): pull}

	for pull.Chain() != 0 {
		link, err := o.repoService.getPullRequest(pull.Chain())

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
