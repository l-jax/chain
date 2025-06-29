package tui

import (
	"chain/chain"
)

type adaptor struct {
	orchestrator *chain.Orchestrator
	targetLabel  string
}

func newAdaptor(targetLabel string) *adaptor {
	return &adaptor{
		orchestrator: chain.InitOrchestrator(targetLabel),
		targetLabel:  targetLabel,
	}
}

func (a *adaptor) ListItems() ([]*Item, error) {
	prs, err := a.orchestrator.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0, len(prs))
	for _, pr := range prs {
		item := newItem(
			pr.Id(),
			pr.Title(),
			pr.Branch(),
			pr.Body(),
			pr.State().String(),
			pr.LinkId(),
			pr.Blocked(),
			pr.HasLabel(a.targetLabel),
		)
		items = append(items, item)
	}

	return items, nil
}

func (a *adaptor) GetItemsLinkedTo(item *Item) ([]*Item, error) {
	prs, err := a.orchestrator.GetPrsLinkedTo(item.Id())
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0, len(prs))
	for _, p := range prs {
		items = append(items, newItem(
			p.Id(),
			p.Title(),
			p.Branch(),
			p.Body(),
			p.State().String(),
			p.LinkId(),
			p.Blocked(),
			p.HasLabel(a.targetLabel),
		))
	}
	return items, nil
}
