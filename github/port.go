package github

import (
	"encoding/json"

	"github.com/cli/go-gh/v2"
)

const jsonFields = "title,body,url,state,labels,mergeable,headRefName,number"

type port struct{}

func (p port) GetPr(number string) (*gitHubPr, error) {
	out, _, err := gh.Exec("pr", "view", number, "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var pr gitHubPr
	err = json.Unmarshal(out.Bytes(), &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func (p port) ListActivePrs() ([]*gitHubPr, error) {
	out, _, err := gh.Exec("pr", "list", "--author", "@me", "--json", jsonFields)
	if err != nil {
		return nil, err
	}

	var prs []*gitHubPr
	err = json.Unmarshal(out.Bytes(), &prs)

	if err != nil {
		return nil, err
	}
	return prs, nil
}
