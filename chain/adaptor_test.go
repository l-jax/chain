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

func FindLinkedPrNumberInBody(t *testing.T) {
	body := "do not merge until #123 is released"
	want := uint(123)

	got := findLinkedPrNumberInBody(body)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMap(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			prMock := givenAMockPr(test.branch, test.ghState, []github.GhLabel{{Name: test.label}}, 0)
			pr, err := mapPr(&prMock)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			assertPrMappedCorrectly(t, pr, prMock, state, 0)
		})
	}
}

func TestMapShouldErrorIfUnexpectedState(t *testing.T) {
	state := "unexpected"
	want := fmt.Sprintf("failed to map pull request some-branch: unexpected state: %s", state)

	mockPr := givenAMockPr("some-branch", state, nil, 0)

	_, err := mapPr(&mockPr)

	assertError(t, err, want)
}

func TestGetPullRequestReturnsMappedPull(t *testing.T) {
	mockPr := givenAMockPr("some-branch", "OPEN", nil, 1)
	mockPrs := []*github.GhPullRequest{&mockPr}

	portMock := newPortMock(mockPrs, false)
	adaptor := ghAdaptor{portMock}

	pull, err := adaptor.getPullRequest(1)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assertPrMappedCorrectly(t, pull, mockPr, StateOpen, 1)
}

func TestGetPullRequestReturnsErrorIfPortErrors(t *testing.T) {
	want := fmt.Sprintf("failed to fetch pull 1: %s", ErrPortMock)

	portMock := newPortMock(nil, true)
	adaptor := ghAdaptor{portMock}

	_, err := adaptor.getPullRequest(1)

	assertError(t, err, want)
}

func TestListPullRequests(t *testing.T) {
	mockPr := givenAMockPr("branch-1", "OPEN", nil, 0)
	mockPr2 := givenAMockPr("branch-2", "MERGED", nil, 0)
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

func givenAMockPr(branch, state string, labels []github.GhLabel, chain uint) github.GhPullRequest {
	mockPr := github.GhPullRequest{
		Title:       "mock",
		Body:        fmt.Sprintf("do not merge until #%d is released", chain),
		HeadRefName: branch,
		Url:         "github.com",
		State:       state,
		Labels:      labels,
		Mergeable:   "true",
		Number:      1,
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

func assertPrMappedCorrectly(t *testing.T, got *Pull, want github.GhPullRequest, state State, chain uint) {
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

	if got.Chain() != chain {
		t.Errorf("got chain %d want %d", got.Chain(), chain)
	}
}
