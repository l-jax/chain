package github

import "fmt"

type gitHubPort interface {
	GetPr(number string) (*gitHubPr, error)
	ListActivePrs() ([]*gitHubPr, error)
}

type Adaptor struct {
	port gitHubPort
}

func NewAdaptor() *Adaptor {
	return &Adaptor{port: &port{}}
}

func (a *Adaptor) GetPr(number uint) (*PullRequest, error) {
	pr, err := a.port.GetPr(fmt.Sprintf("%d", number))
	if err != nil {
		return nil, fmt.Errorf("%w %d: %w", ErrFailedToFetch, number, err)
	}
	return mapPr(pr)
}

func (a *Adaptor) ListOpenPrs() ([]*PullRequest, error) {
	gitHubPrs, err := a.port.ListActivePrs()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	pullRequests := make([]*PullRequest, 0, len(gitHubPrs))
	for _, pr := range gitHubPrs {
		pullResult, err := mapPr(pr)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
		}
		pullRequests = append(pullRequests, pullResult)
	}
	return pullRequests, nil
}

func mapPr(pr *gitHubPr) (*PullRequest, error) {
	state, err := mapState(pr.State)

	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", ErrFailedToMap, pr.HeadRefName, err)
	}

	labels := make([]string, 0, len(pr.Labels))
	for _, label := range pr.Labels {
		labels = append(labels, label.Name)
	}
	return NewPullRequest(pr.Title, pr.HeadRefName, pr.Body, state, labels, pr.Number), nil
}

func mapState(state string) (State, error) {
	switch state {
	case "OPEN":
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		return StateMerged, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnexpectedState, state)
	}
}
