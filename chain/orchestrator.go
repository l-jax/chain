package chain

func GetFakePullRequests() []Pull {
	return []Pull{
		{"my pull request", "some-branch", StateOpen, nil},
		{"some other pr", "some-different-branch", StateReleased, nil},
		{"a change", "my-branch", StateMerged, nil},
		{"code", "some-other-branch", StateOpen, nil},
	}
}
