package chain

import (
	"chain/github"
	"testing"
)

func TestMapPr(t *testing.T) {
	pr := github.NewPullRequest(
		"Add feature",
		"feature-branch",
		"do not merge until #42 is released",
		github.StateOpen,
		[]string{"DO NOT MERGE"},
		41,
	)

	want := NewPr(
		"Add feature",
		"feature-branch",
		"do not merge until #42 is released",
		41,
		42,
		open,
		true,
	)

	got, err := mapGitHubPullRequest(pr, 42, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapState(t *testing.T) {
	tests := []struct {
		name  string
		state github.State
		want  state
	}{
		{"draft", github.StateDraft, draft},
		{"open", github.StateOpen, open},
		{"closed", github.StateClosed, closed},
		{"merged without released label", github.StateMerged, merged}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapState(tt.state, []string{})
			if err != nil {
				t.Fatalf("%s: expected no error, got %v", tt.name, err)
			}
			if got != tt.want {
				t.Errorf("%s: got %q want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestFindLink(t *testing.T) {
	body := "do not merge until #123 is released"
	want := uint(123)

	got := findLinkedPr(body)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFindLinkCaseInsensitive(t *testing.T) {
	body := "do nOt MERGE until #123 is released"
	want := uint(123)

	got := findLinkedPr(body)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
