package chain

import "fmt"

var (
	ErrUnexpectedState = fmt.Errorf("unexpected state")
	ErrFailedToFetch   = fmt.Errorf("failed to fetch")
	ErrFailedToMap     = fmt.Errorf("failed to map pull request")
	ErrLoopedChain     = fmt.Errorf("looped chain detected, please check the pull request links")
)
