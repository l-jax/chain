package chain

import (
	"chain/github"
	"errors"
	"fmt"
	"testing"
)

var stateMappingTests = map[State]struct {
	branch, ghState, label string
}{
	StateOpen:     {branch: "my-open-branch", ghState: "OPEN", label: ""},
	StateBlocked:  {branch: "my-blocked-branch", ghState: "OPEN", label: "DO NoT MeRGE"},
	StateMerged:   {branch: "my-merged-branch", ghState: "MERGED", label: ""},
	StateReleased: {branch: "my-released-branch", ghState: "MERGED", label: "RELEaSED"},
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

func TestMapToPullRequest(t *testing.T) {
	for state, test := range stateMappingTests {
		t.Run(state.String(), func(t *testing.T) {
			pr := givenAGitHubPr(test.branch, test.ghState, []github.Label{{Name: test.label}}, 0)
			got, err := mapToPullRequest(&pr)

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

	_, err := mapToPullRequest(&mockPr)

	assertError(t, err, want)
}

func TestGetPullRequestReturnsMappedPull(t *testing.T) {
	mockPr := givenAGitHubPr("some-branch", "OPEN", nil, 1)
	mockPrs := []*github.PullRequest{&mockPr}

	clientStub := newClientStub(mockPrs, false)
	service := gitHubService{clientStub}

	pull, err := service.getPullRequest(1)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assertPrMappedCorrectly(t, pull, mockPr, StateOpen, 1)
}

func TestGetPullRequestReturnsErrorIfClientErrors(t *testing.T) {
	want := fmt.Errorf("%w 1: %w", ErrFailedToFetch, errClientStub)

	clientStub := newClientStub(nil, true)
	service := gitHubService{clientStub}

	_, err := service.getPullRequest(1)

	assertError(t, err, want)
}

func TestListPullRequests(t *testing.T) {
	mockPr := givenAGitHubPr("branch-1", "OPEN", nil, 0)
	mockPr2 := givenAGitHubPr("branch-2", "MERGED", nil, 0)
	want := []*github.PullRequest{&mockPr, &mockPr2}

	clientStub := newClientStub(want, false)
	service := gitHubService{clientStub}

	got, _ := service.listPullRequests()

	if len(got) != len(want) {
		t.Errorf("got %q pull requests want %q", len(got), len(want))
	}
}

func TestListPullRequestsReturnsErrorIfClientErrors(t *testing.T) {
	clientStub := newClientStub(nil, true)
	service := gitHubService{clientStub}
	want := fmt.Errorf("%w: %w", ErrFailedToFetch, errClientStub)

	_, err := service.listPullRequests()

	assertError(t, err, want)
}

func givenAGitHubPr(branch, state string, labels []github.Label, chain uint) github.PullRequest {
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

var errClientStub = errors.New("mock stub error")

type clientStub struct {
	githubPullRequests []*github.PullRequest
}

func newClientStub(pulls []*github.PullRequest, shouldError bool) *clientStub {
	return &clientStub{githubPullRequests: pulls}
}

func (p *clientStub) GetPr(branch string) (*github.PullRequest, error) {
	if p.githubPullRequests == nil {
		return nil, errClientStub
	}
	return p.githubPullRequests[0], nil
}

func (p *clientStub) ListActivePrs() ([]*github.PullRequest, error) {
	if p.githubPullRequests == nil {
		return nil, errClientStub
	}
	return p.githubPullRequests, nil
}
