package chain

import "testing"

func TestGetPullRequests(t *testing.T) {

	want := []*Pull{
		{"my pull request", "some-branch", "some body", StateOpen, 2, 0},
	}

	adaptor := &AdaptorMock{pulls: want}

	orchestrator := NewOrchestrator(adaptor)

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
	releasedPr := &Pull{"remove something", "some-branch-123", "some body", StateReleased, 1, 0}
	mergedPr := &Pull{"add something", "my-branch", "do not merge until #14 is released", StateOpen, 12, 14}
	openPr := &Pull{"do something", "branch", "some body", StateOpen, 14, 0}

	want := []*Pull{
		mergedPr,
		openPr,
	}

	adaptor := &AdaptorMock{pulls: []*Pull{openPr, mergedPr, releasedPr}}
	orchestrator := NewOrchestrator(adaptor)

	got, err := orchestrator.GetChain(12)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pulls, got %d", len(want), len(got))
	}

	for i, pull := range got {
		if pull != want[i] {
			t.Errorf("expected pull %d to be %v, got %v", i, want[i], pull)
		}
	}
}
