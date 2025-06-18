package chain

import "testing"

func TestGetPullRequests(t *testing.T) {

	want := []*PullRequest{
		{"my pull request", "some-branch", "some body", StateOpen, 2, 0},
	}

	adaptor := &AdaptorMock{pulls: want}
	orchestrator := orchestrator{adaptor: adaptor}

	got, err := orchestrator.GetPullRequests()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pull requests, got %d", len(want), len(got))
	}

	if adaptor.listCalls != 1 {
		t.Fatalf("expected listPullRequests to be called once, got %d", adaptor.listCalls)
	}
}

func TestGetChain(t *testing.T) {
	releasedPr := &PullRequest{"remove something", "some-branch-123", "some body", StateReleased, 1, 0}
	mergedPr := &PullRequest{"add something", "my-branch", "do not merge until #14 is released", StateOpen, 12, 14}
	openPr := &PullRequest{"do something", "branch", "some body", StateOpen, 11, 0}
	blockedPr := &PullRequest{"update something", "another-branch", "do not merge until #11 is released", StateBlocked, 14, 11}

	want := []*PullRequest{
		mergedPr,
		blockedPr,
		openPr,
	}

	adaptor := &AdaptorMock{pulls: []*PullRequest{openPr, mergedPr, releasedPr, blockedPr}}
	orchestrator := orchestrator{adaptor: adaptor}

	got, err := orchestrator.GetChain(12)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pulls, got %d", len(want), len(got))
	}

	for i, link := range want {
		if link != got[link.Number()] {
			t.Errorf("expected pull %d to be %v, got %v", i, want[i], link)
		}
	}
}

func TestGetChainErrorIfLooped(t *testing.T) {
	adaptor := &AdaptorMock{pulls: []*PullRequest{
		{"add something", "my-branch", "do not merge until #11 is released", StateOpen, 12, 11},
		{"do something", "branch", "do not merge until #12 is released", StateOpen, 11, 12},
	}}
	orchestrator := orchestrator{adaptor: adaptor}

	_, err := orchestrator.GetChain(12)

	if err == nil {
		t.Fatalf("expected error")
	}

	if err.Error() != ErrLoopedChain.Error() {
		t.Fatalf("expected error %v, got %v", ErrLoopedChain, err)
	}
}
