package main

type PullRequest struct {
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Url       string  `json:"url"`
	State     string  `json:"state"`
	Labels    []Label `json:"labels"`
	Mergeable string  `json:"mergeable"`
}

type Label struct {
	Name string `json:"name"`
}
