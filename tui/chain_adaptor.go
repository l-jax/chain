package tui

import (
	"chain/chain"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss/tree"
)

func getActivePullRequests() ([]list.Item, error) {
	orchestrator := chain.NewOrchestrator()
	pulls, err := orchestrator.GetPullRequests()

	if err != nil {
		return nil, err //TODO: error

	}

	items := make([]list.Item, len(pulls))
	for i := range pulls {
		items[i] = item{pulls[i].Title(), pulls[i].Branch()}
	}

	return items, nil
}

func getTree(branch string) *tree.Tree {
	pulls := chain.GetChain(branch)
	root := tree.Root(branch)

	_ = addBranches(root, pulls)
	return root
}

func addBranches(root *tree.Tree, pulls []chain.Pull) error {
	for _, pull := range pulls {
		if pull.Branch() == root.Value() && pull.Chain() != 0 {
			treeBranch := tree.Root(pull.Chain())
			root.Child(treeBranch)

			if err := addBranches(treeBranch, pulls); err != nil {
				return err
			}
		}
	}
	return nil
}
