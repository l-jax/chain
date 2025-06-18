package chain

import (
	"chain/github"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrFailedToFetch   = fmt.Errorf("failed to fetch")
	ErrFailedToMap     = fmt.Errorf("failed to map pull request")
	ErrUnexpectedState = fmt.Errorf("unexpected state")
)

type GitHubClient interface {
	GetPr(number string) (*github.PullRequest, error)
	ListActivePrs() ([]*github.PullRequest, error)
}

type gitHubService struct {
	gitHubClient GitHubClient
}

func NewGitHubService(gitHubClient GitHubClient) *gitHubService {
	return &gitHubService{
		gitHubClient: gitHubClient,
	}
}

func (a *gitHubService) getPullRequest(number uint) (*PullRequest, error) {
	pr, err := a.gitHubClient.GetPr(fmt.Sprintf("%d", number))
	if err != nil {
		return nil, fmt.Errorf("%w %d: %w", ErrFailedToFetch, number, err)
	}
	return mapToPullRequest(pr)
}

func (a *gitHubService) listPullRequests() ([]*PullRequest, error) {
	gitHubPrs, err := a.gitHubClient.ListActivePrs()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
	}

	pullRequests := make([]*PullRequest, 0, len(gitHubPrs))
	for _, pr := range gitHubPrs {
		pullResult, err := mapToPullRequest(pr)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrFailedToFetch, err)
		}
		pullRequests = append(pullRequests, pullResult)
	}
	return pullRequests, nil
}

func findLink(body string) uint {
	re := regexp.MustCompile(`do not merge until #(\d+)`)
	match := re.FindStringSubmatch(body)

	if match == nil {
		return 0
	}

	link, err := strconv.ParseUint(match[1], 10, 32)
	if err != nil {
		return 0
	}
	return uint(link)
}

func mapToPullRequest(pr *github.PullRequest) (*PullRequest, error) {
	state, err := mapState(pr.State, pr.Labels)

	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", ErrFailedToMap, pr.HeadRefName, err)
	}

	link := findLink(pr.Body)

	return NewPullRequest(pr.Title, pr.HeadRefName, pr.Body, state, pr.Number, link), nil
}

func mapState(state string, labels []github.Label) (State, error) {
	switch state {
	case "OPEN":
		if isBlocked(labels) {
			return StateBlocked, nil
		}
		return StateOpen, nil
	case "CLOSED":
		return StateClosed, nil
	case "MERGED":
		if isReleased(labels) {
			return StateReleased, nil
		}
		return StateMerged, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnexpectedState, state)
	}
}

func isBlocked(labels []github.Label) bool {
	for _, label := range labels {
		if label.Name == "DO NOT MERGE" {
			return true
		}
	}
	return false
}

func isReleased(labels []github.Label) bool {
	for _, label := range labels {
		if label.Name == "RELEASED" {
			return true
		}
	}
	return false
}
