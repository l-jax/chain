package chain

import (
	"chain/github"
	"errors"
)

var ErrPortMock = errors.New("mock port error")

type PortMock struct {
	pulls       []*github.GhPullRequest
	shouldError bool
}

func newPortMock(pulls []*github.GhPullRequest, shouldError bool) *PortMock {
	return &PortMock{pulls: pulls, shouldError: shouldError}
}

func (p *PortMock) GetPr(branch string) (*github.GhPullRequest, error) {
	if p.shouldError {
		return nil, ErrPortMock
	}
	return p.pulls[0], nil
}

func (p *PortMock) ListActivePrs() ([]*github.GhPullRequest, error) {
	if p.shouldError {
		return nil, ErrPortMock
	}
	return p.pulls, nil
}

type AdaptorMock struct {
	pulls     []*Pull
	listCalls int
	getCalls  int
}

func (a *AdaptorMock) listPullRequests() ([]*Pull, error) {
	a.listCalls++
	return a.pulls, nil
}

func (a *AdaptorMock) getPullRequest(branch string) (*Pull, error) {
	a.getCalls++
	return a.pulls[0], nil
}
