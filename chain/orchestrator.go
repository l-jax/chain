package chain

type orchestrator struct {
	adaptor adaptor
}

func NewOrchestrator(adaptor adaptor) *orchestrator {
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

	chain := make([]*Pull, 0)
	chain = append(chain, pull)

	if err != nil {
		return nil, err
	}

	linkedPull, err := o.adaptor.getPullRequest(pull.Chain())

	if err != nil {
		return nil, err
	}

	chain = append(chain, linkedPull)

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
