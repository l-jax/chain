package github

import (
	"errors"
	"fmt"
	"testing"
)

var stateMappingTests = map[State]struct {
	branch, ghState, label string
}{
	StateOpen:   {branch: "my-open-branch", ghState: "OPEN", label: "DO NOT MERGE"},
	StateMerged: {branch: "my-merged-branch", ghState: "MERGED", label: "RELEASED"},
	StateClosed: {branch: "my-closed-branch", ghState: "CLOSED", label: ""},
}

func TestMapToPullRequest(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			pr := givenAGitHubPr(test.branch, test.ghState, []gitHubLabel{{Name: test.label}}, 0)
			got, err := mapPr(&pr)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			assertPrMappedCorrectly(t, got, pr, state, 0)
		})
	}
}

func TestMapToPullShouldErrorIfUnexpectedState(t *testing.T) {
	state := "unexpected"

	mockPr := givenAGitHubPr("some-branch", state, nil, 0)
	want := fmt.Errorf("%w some-branch: %w: %s", ErrFailedToMap, ErrUnexpectedState, state)

	_, err := mapPr(&mockPr)

	assertError(t, err, want)
}

func TestGetPullRequestReturnsMappedPull(t *testing.T) {
	mockPr := givenAGitHubPr("some-branch", "OPEN", nil, 1)
	mockPrs := []*gitHubPr{&mockPr}

	clientStub := initPortStub(mockPrs, false)
	service := Adaptor{port: clientStub}

	pull, err := service.GetPr(1)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assertPrMappedCorrectly(t, pull, mockPr, StateOpen, 1)
}

func TestGetPullRequestReturnsErrorIfClientErrors(t *testing.T) {
	want := fmt.Errorf("%w 1: %w", ErrFailedToFetch, errPortStub)

	clientStub := initPortStub(nil, true)
	service := Adaptor{port: clientStub}

	_, err := service.GetPr(1)

	assertError(t, err, want)
}

func TestListPullRequests(t *testing.T) {
	mockPr := givenAGitHubPr("branch-1", "OPEN", nil, 0)
	mockPr2 := givenAGitHubPr("branch-2", "MERGED", nil, 0)
	want := []*gitHubPr{&mockPr, &mockPr2}

	clientStub := initPortStub(want, false)
	service := Adaptor{port: clientStub}

	got, _ := service.ListOpenPrs()

	if len(got) != len(want) {
		t.Errorf("got %q pull requests want %q", len(got), len(want))
	}
}

func TestListPullRequestsReturnsErrorIfClientErrors(t *testing.T) {
	clientStub := initPortStub(nil, true)
	service := Adaptor{port: clientStub}
	want := fmt.Errorf("%w: %w", ErrFailedToFetch, errPortStub)

	_, err := service.ListOpenPrs()

	assertError(t, err, want)
}

func givenAGitHubPr(branch, state string, labels []gitHubLabel, chain uint) gitHubPr {
	mockPr := gitHubPr{
		Title:       "mock",
		Body:        fmt.Sprintf("do not merge until #%d is released", chain),
		HeadRefName: branch,
		State:       state,
		Labels:      labels,
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

func assertPrMappedCorrectly(t *testing.T, got *PullRequest, want gitHubPr, state State, chain uint) {
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
}

var errPortStub = errors.New("stub error")

type portStub struct {
	githubPullRequests []*gitHubPr
}

func initPortStub(pulls []*gitHubPr, shouldError bool) *portStub {
	return &portStub{githubPullRequests: pulls}
}

func (p *portStub) GetPr(branch string) (*gitHubPr, error) {
	if p.githubPullRequests == nil {
		return nil, errPortStub
	}
	return p.githubPullRequests[0], nil
}

func (p *portStub) ListActivePrs() ([]*gitHubPr, error) {
	if p.githubPullRequests == nil {
		return nil, errPortStub
	}
	return p.githubPullRequests, nil
}
