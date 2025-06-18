package chain

import (
	"chain/github"
	"fmt"
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

func TestFindLink(t *testing.T) {
	body := "do not merge until #123 is released"
	want := uint(123)

	got := findLink(body)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMapToPull(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			prMock := givenAMockPr(test.branch, test.ghState, []github.Label{{Name: test.label}}, 0)
			pr, err := mapToPull(&prMock)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			assertPrMappedCorrectly(t, pr, prMock, state, 0)
		})
	}
}

func TestMapToPullShouldErrorIfUnexpectedState(t *testing.T) {
	state := "unexpected"

	mockPr := givenAMockPr("some-branch", state, nil, 0)
	want := fmt.Errorf("%w some-branch: %w: %s", ErrFailedToMap, ErrUnexpectedState, state)

	_, err := mapToPull(&mockPr)

	assertError(t, err, want)
}

func TestGetPullRequestReturnsMappedPull(t *testing.T) {
	mockPr := givenAMockPr("some-branch", "OPEN", nil, 1)
	mockPrs := []*github.PullRequest{&mockPr}

	portMock := newPortMock(mockPrs, false)
	adaptor := gitHubAdaptor{portMock}

	pull, err := adaptor.getPullRequest(1)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assertPrMappedCorrectly(t, pull, mockPr, StateOpen, 1)
}

func TestGetPullRequestReturnsErrorIfPortErrors(t *testing.T) {
	want := fmt.Errorf("%w 1: %w", ErrFailedToFetch, ErrPortMock)

	portMock := newPortMock(nil, true)
	adaptor := gitHubAdaptor{portMock}

	_, err := adaptor.getPullRequest(1)

	assertError(t, err, want)
}

func TestListPullRequests(t *testing.T) {
	mockPr := givenAMockPr("branch-1", "OPEN", nil, 0)
	mockPr2 := givenAMockPr("branch-2", "MERGED", nil, 0)
	mockPrs := []*github.PullRequest{&mockPr, &mockPr2}

	mockPort := newPortMock(mockPrs, false)
	adaptor := gitHubAdaptor{mockPort}

	pulls, _ := adaptor.listPullRequests()

	if len(pulls) != len(mockPrs) {
		t.Errorf("got %q pull requests want %q", len(pulls), len(mockPrs))
	}
}

func TestListPullRequestsReturnsErrorIfPortErrors(t *testing.T) {
	portMock := newPortMock(nil, true)
	adaptor := gitHubAdaptor{portMock}
	want := fmt.Errorf("%w: %w", ErrFailedToFetch, ErrPortMock)

	_, err := adaptor.listPullRequests()

	assertError(t, err, want)
}

func givenAMockPr(branch, state string, labels []github.Label, chain uint) github.PullRequest {
	mockPr := github.PullRequest{
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

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Error("expected error")
		return
	}

	if got.Error() != want.Error() {
		if unwrapped := fmt.Errorf("%w", got); unwrapped.Error() == want.Error() {
			return
		}
		t.Errorf("got error %s, want %s", got.Error(), want.Error())
	}
}

func assertPrMappedCorrectly(t *testing.T, got *PullRequest, want github.PullRequest, state State, chain uint) {
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
