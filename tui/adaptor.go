package tui

import "chain/chain"

type adaptor struct {
	orchestrator *chain.Orchestrator
	items        []*Item
	linkedItems  map[uint][]*Item
}

func newAdaptor() *adaptor {
	return &adaptor{
		orchestrator: chain.InitOrchestrator("RELEASED"),
		items:        []*Item{},
		linkedItems:  make(map[uint][]*Item),
	}
}

func (h *adaptor) ListItems(refresh bool) ([]*Item, error) {
	if !refresh && len(h.items) > 0 {
		return h.items, nil
	}

	prs, err := h.orchestrator.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	h.items = make([]*Item, 0, len(prs))
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
		h.items = append(h.items, item)
	}

	return h.items, nil
}

func (h *adaptor) GetItemsLinkedTo(item *Item, refresh bool) ([]*Item, error) {
	if !refresh && h.linkedItems[item.Id()] != nil {
		return h.linkedItems[item.Id()], nil
	}

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

	h.linkedItems[item.Id()] = items
	return items, nil
}
