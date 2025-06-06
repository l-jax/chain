package main

import (
	"encoding/json"
	"github.com/cli/go-gh/v2"
)

const jsonFields = "title,body,url,state,labels,mergeable"

func getPr(url string) (*PullRequest, error) {
	out, _, err := gh.Exec("pr", "view", url, "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var pr PullRequest
	err = json.Unmarshal(out.Bytes(), &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func listPrs(label, state string) ([]PullRequest, error) {
	out, _, err := gh.Exec("pr", "list", "--author", "@me", "--label", label, "--state", state, "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var prs []PullRequest
	err = json.Unmarshal(out.Bytes(), &prs)

	if err != nil {
		return nil, err
	}
	return prs, nil
}

func viewPr(url string) error {
	_, _, err := gh.Exec("pr", "view", url, "--web")
	if err != nil {
		return err
	}
	return nil
}

func addLabel(url, label string) error {
	_, _, err := gh.Exec("pr", "edit", url, "--add-label", label)
	if err != nil {
		return err
	}
	return nil
}

func removeLabel(url, label string) error {
	_, _, err := gh.Exec("pr", "edit", url, "--remove-label", label)
	if err != nil {
		return err
	}
	return nil
}
