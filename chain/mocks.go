package chain

import (
	"chain/github"
	"errors"
)

var ErrPortMock = errors.New("mock port error")

type PortMock struct {
	pulls       []*github.PullRequest
	shouldError bool
}

func newPortMock(pulls []*github.PullRequest, shouldError bool) *PortMock {
	return &PortMock{pulls: pulls, shouldError: shouldError}
}

func (p *PortMock) GetPr(branch string) (*github.PullRequest, error) {
	if p.shouldError {
		return nil, ErrPortMock
	}
	return p.pulls[0], nil
}

func (p *PortMock) ListActivePrs() ([]*github.PullRequest, error) {
	if p.shouldError {
		return nil, ErrPortMock
	}
	return p.pulls, nil
}

type AdaptorMock struct {
	pulls     []*PullRequest
	listCalls int
	getCalls  int
}

func (a *AdaptorMock) listPullRequests() ([]*PullRequest, error) {
	a.listCalls++
	return a.pulls, nil
}

func (a *AdaptorMock) getPullRequest(number uint) (*PullRequest, error) {
	a.getCalls++
	for _, pull := range a.pulls {
		if pull.Number() == number {
			return pull, nil
		}

	}
	return nil, nil
}
