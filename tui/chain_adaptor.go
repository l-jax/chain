package tui

import (
	"chain/chain"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss/tree"
	"slices"
)

func getActivePullRequests() []list.Item {
	pulls := chain.GetOpen()
	active := slices.Collect(func(yield func(pull chain.Pull) bool) {
		for _, p := range pulls {
			if p.State() == chain.StateOpen {
				if !yield(p) {
					return
				}
			}
		}
	})

	items := make([]list.Item, len(active))
	for i := range active {
		items[i] = item{active[i].Title(), active[i].Branch()}
	}
	return items
}

func getTree(branch string) *tree.Tree {
	pulls := chain.GetChain(branch)
	root := tree.Root(branch)

	_ = addBranches(root, pulls)
	return root
}

func addBranches(root *tree.Tree, pulls []chain.Pull) error {
	for _, pull := range pulls {
		if pull.Branch() == root.Value() && pull.Chain() != nil {
			treeBranch := tree.Root(pull.Chain().Branch())
			root.Child(treeBranch)

			if err := addBranches(treeBranch, pulls); err != nil {
				return err
			}
		}
	}
	return nil
}
