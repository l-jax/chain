package github

import "fmt"

var (
	ErrFailedToFetch   = fmt.Errorf("failed to fetch")
	ErrFailedToMap     = fmt.Errorf("failed to map pull request")
	ErrUnexpectedState = fmt.Errorf("unexpected state")
)
