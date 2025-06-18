package tui

import (
	"chain/chain"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss/tree"
)

func getActivePullRequests() ([]list.Item, error) {
	handler := chain.NewChainHandler()
	pulls, err := handler.GetPullRequests()

	if err != nil {
		return nil, err //TODO: error

	}

	items := make([]list.Item, len(pulls))
	for i := range pulls {
		items[i] = item{pulls[i].Title(), pulls[i].Branch()}
	}

	return items, nil
}

func getTree(root uint) (*tree.Tree, error) {
	handler := chain.NewChainHandler()
	chain, err := handler.GetChain(root)

	if err != nil {
		return nil, err //TODO: error

	}

	tree, err := buildTree(root, chain)
	if err != nil {
		return nil, err //TODO: error

	}

	return tree, nil
}

func buildTree(root uint, chain map[uint]*chain.PullRequest) (*tree.Tree, error) {
	if len(chain) == 0 {
		return nil, nil // Nothing to add if chain is nil
	}

	tree := tree.Root(chain[root].Branch())

	node := root
	for node != 0 {
		link, exists := chain[node]
		if !exists {
			return nil, fmt.Errorf("linked pull request with number %d not found in chain", node)
		}
		tree.SetValue(link.Branch())
		node = link.Chain()
	}

	return tree, nil
}
