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
