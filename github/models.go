package github

type gitHubPr struct {
	Title       string        `json:"title"`
	Body        string        `json:"body"`
	HeadRefName string        `json:"headRefName"`
	Url         string        `json:"url"`
	State       string        `json:"state"`
	Number      uint          `json:"number"`
	Labels      []gitHubLabel `json:"labels"`
	Mergeable   string        `json:"mergeable"`
}

type gitHubLabel struct {
	Name string `json:"name"`
}
