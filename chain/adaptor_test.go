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
	StateOpen:     {branch: "my-open-branch", ghState: "OPEN", label: ""},
	StateBlocked:  {branch: "my-blocked-branch", ghState: "OPEN", label: "DO NOT MERGE"},
	StateMerged:   {branch: "my-merged-branch", ghState: "MERGED", label: ""},
	StateReleased: {branch: "my-released-branch", ghState: "MERGED", label: "RELEASED"},
	StateClosed:   {branch: "my-closed-branch", ghState: "CLOSED", label: "DO NOT MERGE"},
}

func TestMap(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			prMock := givenAMockPr(test.branch, test.ghState, []github.GhLabel{{Name: test.label}})
			pr, err := mapPr(&prMock)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			assertPrMappedCorrectly(t, pr, prMock, state)
		})
	}
}

func TestGetPullRequestReturnsMappedPull(t *testing.T) {
	branch := "branch-name"
	mockPr := givenAMockPr(branch, "OPEN", nil)
	mockPrs := []*github.GhPullRequest{&mockPr}

	portMock := newPortMock(mockPrs, false)
	adaptor := ghAdaptor{portMock}

	pull, err := adaptor.getPullRequest(branch)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assertPrMappedCorrectly(t, pull, mockPr, StateOpen)
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

func TestListPullRequests(t *testing.T) {
	mockPr := givenAMockPr("branch-1", "OPEN", nil)
	mockPr2 := givenAMockPr("branch-2", "MERGED", nil)
	mockPrs := []*github.GhPullRequest{&mockPr, &mockPr2}

	mockPort := newPortMock(mockPrs, false)
	adaptor := ghAdaptor{mockPort}

	pulls, _ := adaptor.listPullRequests()

	if len(pulls) != len(mockPrs) {
		t.Errorf("got %q pull requests want %q", len(pulls), len(mockPrs))
	}
}

func TestListPullRequestsReturnsErrorIfPortErrors(t *testing.T) {
	portMock := newPortMock(nil, true)
	adaptor := ghAdaptor{portMock}
	want := "failed to fetch all: " + ErrPortMock.Error()

	_, err := adaptor.listPullRequests()

	assertError(t, err, want)
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

	if got.Body() != want.Body {
		t.Errorf("got body %s want %q", got.body, want.Body)
	}

	if got.State() != state {
		t.Errorf("got state %q want %q", got.State(), state)
	}

	if got.Chain() != nil {
		t.Errorf("got chain %s want nil", got.Chain().branch)
	}
}
