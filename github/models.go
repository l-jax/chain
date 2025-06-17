package github

type GhPullRequest struct {
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	HeadRefName string    `json:"headRefName"`
	Url         string    `json:"url"`
	State       string    `json:"state"`
	Number      uint      `json:"number"`
	Labels      []GhLabel `json:"labels"`
	Mergeable   string    `json:"mergeable"`
}

type GhLabel struct {
	Name string `json:"name"`
}
