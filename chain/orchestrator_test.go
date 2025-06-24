package chain

import (
	"chain/github"
	"testing"
)

func TestListOpenPrs(t *testing.T) {

	want := []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #14 is released", github.StateOpen, []string{}, 12),
	}

	service := &serviceFake{pullRequests: want}
	handler := orchestrator{gitHubAdaptor: service}

	got, err := handler.ListOpenPrs()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pull requests, got %d", len(want), len(got))
	}
}

func TestGetChain(t *testing.T) {
	releasedPr := github.NewPullRequest("release something", "release-branch", "this is released", github.StateMerged, []string{"RELEASED"}, 1)
	mergedPr := github.NewPullRequest("merge something", "my-branch", "do not merge until #14 is released", github.StateMerged, []string{}, 12)
	openPr := github.NewPullRequest("add something", "my-branch", "message", github.StateOpen, []string{}, 11)
	blockedPr := github.NewPullRequest("do something", "branch", "do not merge until #11 is released", github.StateOpen, []string{"DO NOT MERGE"}, 14)

	want := []*github.PullRequest{
		mergedPr,
		blockedPr,
		openPr,
	}

	service := &serviceFake{pullRequests: []*github.PullRequest{openPr, mergedPr, releasedPr, blockedPr}}
	handler := orchestrator{gitHubAdaptor: service}

	got, err := handler.GetPrsLinkedTo(12)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pulls, got %d", len(want), len(got))
	}

	for i, link := range want {

		if link.Number() != got[link.Number()].Id() {
			t.Errorf("expected pull %d to be %v, got %v", i, want[i], link)
		}
	}
}

func TestGetChainErrorIfLooped(t *testing.T) {
	service := &serviceFake{pullRequests: []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #11 is released", github.StateOpen, []string{}, 12),
		github.NewPullRequest("merge something", "my-branch", "do not merge until #12 is released", github.StateMerged, []string{}, 11),
	}}
	handler := orchestrator{gitHubAdaptor: service}

	_, err := handler.GetPrsLinkedTo(12)

	if err == nil {
		t.Fatalf("expected error")
	}

	if err.Error() != ErrLoopedChain.Error() {
		t.Fatalf("expected error %v, got %v", ErrLoopedChain, err)
	}
}

type serviceFake struct {
	pullRequests []*github.PullRequest
}

func (a *serviceFake) ListOpenPrs() ([]*github.PullRequest, error) {
	return a.pullRequests, nil
}

func (a *serviceFake) GetPr(number uint) (*github.PullRequest, error) {
	for _, pull := range a.pullRequests {
		if pull.Number() == number {
			return pull, nil
		}

	}
	return nil, nil
}
