package github

type PullRequest struct {
	Title       string  `json:"title"`
	Body        string  `json:"body"`
	HeadRefName string  `json:"headRefName"`
	Url         string  `json:"url"`
	State       string  `json:"state"`
	Number      uint    `json:"number"`
	Labels      []Label `json:"labels"`
	Mergeable   string  `json:"mergeable"`
}

type Label struct {
	Name string `json:"name"`
}
