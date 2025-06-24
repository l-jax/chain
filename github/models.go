package github

type gitHubPr struct {
	Title       string        `json:"title"`
	Body        string        `json:"body"`
	HeadRefName string        `json:"headRefName"`
	State       string        `json:"state"`
	Number      uint          `json:"number"`
	Labels      []gitHubLabel `json:"labels"`
}

type gitHubLabel struct {
	Name string `json:"name"`
}
