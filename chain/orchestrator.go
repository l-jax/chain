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

func GetOpen() []Pull {
	// TODO: open logic
	return []Pull{
		{"my pull request", "some-branch", "some body", StateOpen, nil},
		{"code", "some-other-branch", "some body", StateOpen, nil},
	}
}

func GetChain(branch string) []Pull {
	// TODO: chain logic
	three := Pull{"remove something", "some-branch-123", "some body", StateReleased, nil}
	two := Pull{"add something", "my-branch", "some body", StateMerged, &three}
	one := Pull{"do something", branch, "some body", StateOpen, &two}
	return []Pull{
		one,
		two,
		three,
	}
}
