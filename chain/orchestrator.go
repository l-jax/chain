package chain

func GetOpen() []Pull {
	// TODO: open logic
	return []Pull{
		{"my pull request", "some-branch", StateOpen, nil},
		{"code", "some-other-branch", StateOpen, nil},
	}
}

func GetChain(branch string) []Pull {
	// TODO: chain logic
	three := Pull{"remove something", "some-branch-123", StateReleased, nil}
	two := Pull{"add something", "my-branch", StateMerged, &three}
	one := Pull{"do something", branch, StateOpen, &two}
	return []Pull{
		one,
		two,
		three,
	}
}
