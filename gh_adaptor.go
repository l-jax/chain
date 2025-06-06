package main

func getPullRequests() ([]PullRequest, error) {
	prs, err := listActivePrs()
	if err != nil {
		return nil, err
	}

	var pullRequests []PullRequest
	for _, pr := range prs {
		pullRequests = append(pullRequests,
			PullRequest{
				pr.Title,
				pr.HeadRefName,
				[]Condition{
					&StateCondition{pr.State, "OPEN"},
				},
			})
	}
	return pullRequests, nil
}
