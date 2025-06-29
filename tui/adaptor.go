package tui

import "chain/chain"

type adaptor struct {
	orchestrator *chain.Orchestrator
}

func newAdaptor(targetLabel string) *adaptor {
	return &adaptor{
		orchestrator: chain.InitOrchestrator(targetLabel),
	}
}

func (h *adaptor) ListItems() ([]*Item, error) {
	prs, err := h.orchestrator.ListOpenPrs()
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
		)
		items = append(items, item)
	}

	return items, nil
}

func (h *adaptor) GetItemsLinkedTo(item *Item) ([]*Item, error) {
	prs, err := h.orchestrator.GetPrsLinkedTo(item.Id())
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
		))
	}
	return items, nil
}
