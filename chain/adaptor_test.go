package chain

import (
	"chain/github"
	"fmt"
	"strings"
	"testing"
)

var stateMappingTests = map[State]struct {
	branch, ghState, label string
}{
	StateOpen:     {branch: "my-open-branch", ghState: "OPEN", label: "DO NOT MERGE"},
	StateMerged:   {branch: "my-merged-branch", ghState: "MERGED", label: ""},
	StateReleased: {branch: "my-released-branch", ghState: "MERGED", label: "RELEASED"},
	StateClosed:   {branch: "my-closed-branch", ghState: "CLOSED", label: "DO NOT MERGE"},
}

func TestGetPullRequestMapsCorrectly(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			prMock := givenAMockPr(test.branch, test.ghState, []github.GhLabel{{test.label}})
			mockPrs := []*github.GhPullRequest{&prMock}
			mockPort := newPortMock(mockPrs, false)

			adaptor := ghAdaptor{mockPort}

			pr, err := adaptor.getPullRequest(test.branch)

			if err != nil {
				t.Fatalf("getPullRequest returned unexpected error: %s", err)
			}

			assertPrMappedCorrectly(t, pr, prMock, state)
		})
	}
}

func TestGetPullRequestReturnsErrorIfPortErrors(t *testing.T) {
	branch := "branch-name"
	want := fmt.Sprintf("failed to fetch branch %s: %s", branch, ErrPortMock)

	portMock := newPortMock(nil, true)
	adaptor := ghAdaptor{portMock}

	_, err := adaptor.getPullRequest(branch)

	assertError(t, err, want)
}

func TestGetPullRequestReturnsErrorIfUnexpectedState(t *testing.T) {
	branch := "branch-name"
	state := "unexpected"
	want := fmt.Sprintf("failed to map pull request %s: unexpected state: %s", branch, state)

	mockPr := givenAMockPr(branch, state, nil)
	mockPrs := []*github.GhPullRequest{&mockPr}
	mockPort := newPortMock(mockPrs, false)

	adaptor := ghAdaptor{mockPort}

	_, err := adaptor.getPullRequest(branch)

	assertError(t, err, want)
}

func TestGetPullRequestsMapsCorrectly(t *testing.T) {
	branch := "branch-name"
	state := "OPEN"

	mockPr := givenAMockPr(branch, state, nil)
	mockPrs := []*github.GhPullRequest{&mockPr}

	mockPort := newPortMock(mockPrs, false)
	adaptor := ghAdaptor{mockPort}

	pulls, _ := adaptor.listPullRequests()

	if len(pulls) != len(mockPrs) {
		t.Errorf("got %q pull requests want %q", len(pulls), len(mockPrs))
	}

	assertPrMappedCorrectly(t, pulls[0], mockPr, StateOpen)
}

func givenAMockPr(branch, state string, labels []github.GhLabel) github.GhPullRequest {
	mockPr := github.GhPullRequest{
		Title:       "mock",
		Body:        "body",
		HeadRefName: branch,
		Url:         "github.com",
		State:       state,
		Labels:      labels,
		Mergeable:   "true",
	}
	return mockPr
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()

	if got == nil {
		t.Error("expected error")
		return
	}

	if strings.Compare(got.Error(), want) != 0 {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertPrMappedCorrectly(t *testing.T, got *Pull, want github.GhPullRequest, state State) {
	t.Helper()
	if got.Title() != want.Title {
		t.Errorf("got title %q want %q", got.title, want.Title)
	}

	if got.Branch() != want.HeadRefName {
		t.Errorf("got branch %q want %q", got.branch, want.HeadRefName)
	}

	if got.State() != state {
		t.Errorf("got state %q want %q", got.State(), state)
	}

	if got.Chain() != nil {
		t.Errorf("got chain %s want nil", got.Chain().branch)
	}
}
