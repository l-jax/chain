package chain

import (
	"chain/github"
	"testing"
)

func TestGetPullRequests(t *testing.T) {

	want := []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #14 is released", github.StateOpen, 12, 14),
	}

	service := &serviceFake{pullRequests: want}
	handler := chainHandler{repoService: service}

	got, err := handler.GetPullRequests()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d pull requests, got %d", len(want), len(got))
	}
}

func TestGetChain(t *testing.T) {
	releasedPr := github.NewPullRequest("release something", "release-branch", "this is released", github.StateReleased, 1, 0)
	mergedPr := github.NewPullRequest("merge something", "my-branch", "do not merge until #14 is released", github.StateOpen, 12, 14)
	openPr := github.NewPullRequest("add something", "my-branch", "message", github.StateOpen, 11, 0)
	blockedPr := github.NewPullRequest("do something", "branch", "do not merge until #11 is released", github.StateBlocked, 14, 11)

	want := []*github.PullRequest{
		mergedPr,
		blockedPr,
		openPr,
	}

	service := &serviceFake{pullRequests: []*github.PullRequest{openPr, mergedPr, releasedPr, blockedPr}}
	handler := chainHandler{repoService: service}

	got, err := handler.GetChain(12)

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
	service := &serviceFake{pullRequests: []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #11 is released", github.StateOpen, 12, 11),
		github.NewPullRequest("merge something", "my-branch", "do not merge until #12 is released", github.StateMerged, 11, 12),
	}}
	handler := chainHandler{repoService: service}

	_, err := handler.GetChain(12)

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

func (a *serviceFake) ListPullRequests() ([]*github.PullRequest, error) {
	return a.pullRequests, nil
}

func (a *serviceFake) GetPullRequest(number uint) (*github.PullRequest, error) {
	for _, pull := range a.pullRequests {
		if pull.Number() == number {
			return pull, nil
		}

	}
	return nil, nil
}
