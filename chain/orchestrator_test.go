package chain

import (
	"chain/github"
	"testing"
)

const targetLabel = "RELEASED"

func TestListOpenPrs(t *testing.T) {
	prs := []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #14 is released", github.StateOpen, []string{}, 12),
		github.NewPullRequest("do something", "branch", "some description", github.StateOpen, []string{targetLabel}, 14),
	}

	orchestrator := Orchestrator{
		gitHubAdaptor: &serviceFake{prs: prs},
		targetLabel:   targetLabel,
		prs:           make(map[uint]*Pr),
	}

	got, err := orchestrator.ListOpenPrs()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(prs) {
		t.Fatalf("expected %d pull requests, got %d", len(prs), len(got))
	}
}

func TestGetPrsLinkedTo(t *testing.T) {
	unrelatedPr := github.NewPullRequest("", "", "", github.StateMerged, []string{"RELEASED"}, 1)

	linkedPrs := []*github.PullRequest{
		github.NewPullRequest("", "", "do not merge until #14 is released", github.StateMerged, []string{}, 12),
		github.NewPullRequest("", "", "", github.StateOpen, []string{}, 11),
		github.NewPullRequest("", "", "do not merge until #11 is released", github.StateOpen, []string{"DO NOT MERGE"}, 14),
	}

	orchestrator := Orchestrator{
		gitHubAdaptor: &serviceFake{prs: []*github.PullRequest{unrelatedPr, linkedPrs[0], linkedPrs[1], linkedPrs[2]}},
		targetLabel:   targetLabel,
		prs:           make(map[uint]*Pr),
	}

	got, err := orchestrator.GetPrsLinkedTo(12)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(linkedPrs) {
		t.Fatalf("expected %d pulls, got %d", len(linkedPrs), len(got))
	}

	for i, pr := range linkedPrs {
		if pr.Number() != got[pr.Number()].Id() {
			t.Errorf("expected pull %d to be %v, got %v", i, linkedPrs[i], pr)
		}
	}
}

func TestGetChainErrorIfLooped(t *testing.T) {
	prs := []*github.PullRequest{
		github.NewPullRequest("add something", "my-branch", "do not merge until #11 is released", github.StateOpen, []string{}, 12),
		github.NewPullRequest("merge something", "my-branch", "do not merge until #12 is released", github.StateMerged, []string{}, 11),
	}

	orchestrator := Orchestrator{
		gitHubAdaptor: &serviceFake{prs: prs},
		targetLabel:   targetLabel,
		prs:           make(map[uint]*Pr),
	}

	_, err := orchestrator.GetPrsLinkedTo(12)

	if err == nil {
		t.Fatalf("expected error")
	}

	if err.Error() != ErrLoopedChain.Error() {
		t.Fatalf("expected error %v, got %v", ErrLoopedChain, err)
	}
}

var linkRetrievalTests = map[string]struct {
	label          string
	hasTargetLabel bool
}{
	"Blocked":     {label: "", hasTargetLabel: false},
	"Not blocked": {label: "RELEASED", hasTargetLabel: true},
}

func TestLinkRetrieval(t *testing.T) {
	for name, test := range linkRetrievalTests {
		t.Run(name, func(t *testing.T) {
			pr := github.NewPullRequest("add something", "my-branch", "do not merge until #14 is released", github.StateOpen, []string{}, 12)
			linkedPr := github.NewPullRequest("do something", "branch", "some description", github.StateMerged, []string{test.label}, 14)

			orchestrator := Orchestrator{
				gitHubAdaptor: &serviceFake{prs: []*github.PullRequest{pr, linkedPr}},
				targetLabel:   targetLabel,
				prs:           make(map[uint]*Pr),
			}

			link, err := orchestrator.getLink(linkedPr.Number())

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if link.id != linkedPr.Number() {
				t.Fatalf("expected linked PR id %d, got %d", linkedPr.Number(), link.id)
			}

			if link.hasTargetLabel != test.hasTargetLabel {
				t.Fatalf("expected blocked to be %v, got %v", test.hasTargetLabel, link.hasTargetLabel)
			}
		})
	}
}

type serviceFake struct {
	prs []*github.PullRequest
}

func (a *serviceFake) ListOpenPrs() ([]*github.PullRequest, error) {
	return a.prs, nil
}

func (a *serviceFake) GetPr(number uint) (*github.PullRequest, error) {
	for _, pull := range a.prs {
		if pull.Number() == number {
			return pull, nil
		}

	}
	return nil, nil
}
