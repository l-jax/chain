package tui

import (
	"chain/chain"
	"github.com/charmbracelet/bubbles/list"
	"slices"
)

func getActivePullRequests() []list.Item {
	pulls := chain.GetFakePullRequests()
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
