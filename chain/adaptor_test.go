package chain

import (
	"chain/github"
	"testing"
)

func TestGetPullRequests(t *testing.T) {
	mockPr := github.GhPullRequest{
		Title:       "mock",
		Body:        "body",
		HeadRefName: "branch-name",
		Url:         "github.com",
		State:       "OPEN",
		Mergeable:   "true",
	}

	mockPulls := []github.GhPullRequest{mockPr}

	mockPort := newMockPort(mockPulls)
	adaptor := ghAdaptor{mockPort}

	got, _ := adaptor.getPullRequests()

	if len(got) != len(mockPulls) {
		t.Errorf("got %q pull requests want %q", len(got), len(mockPulls))
	}
}

type MockPort struct {
	pulls []github.GhPullRequest
}

func newMockPort(pulls []github.GhPullRequest) *MockPort {
	return &MockPort{pulls: pulls}
}

func (p *MockPort) ListPrs(_ string) ([]github.GhPullRequest, error) {
	return p.pulls, nil
}
