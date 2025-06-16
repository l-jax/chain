package chain

import "testing"

func TestGetPullRequests(t *testing.T) {

	want := []*Pull{
		{"my pull request", "some-branch", "some body", StateOpen, nil},
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
