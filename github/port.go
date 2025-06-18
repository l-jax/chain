package github

import (
	"encoding/json"

	"github.com/cli/go-gh/v2"
)

const jsonFields = "title,body,url,state,labels,mergeable,headRefName,number"

type Port interface {
	GetPr(number string) (*GhPullRequest, error)
	ListActivePrs() ([]*GhPullRequest, error)
}

type GhPort struct {
}

func (p GhPort) GetPr(number string) (*GhPullRequest, error) {
	out, _, err := gh.Exec("pr", "view", number, "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var pr GhPullRequest
	err = json.Unmarshal(out.Bytes(), &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func (p GhPort) ListActivePrs() ([]*GhPullRequest, error) {
	out, _, err := gh.Exec("pr", "list", "--author", "@me", "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var prs []*GhPullRequest
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
