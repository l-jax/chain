package tui

import (
	"chain/chain"
)

type chainAdaptor struct {
	orchestrator *chain.Orchestrator
	items        []*Item
	linkedItems  map[uint][]*Item
}

func initChainAdaptor() *chainAdaptor {
	return &chainAdaptor{
		orchestrator: chain.InitOrchestrator(),
		items:        []*Item{},
		linkedItems:  make(map[uint][]*Item),
	}
}

func (h *chainAdaptor) ListItems(refresh bool) ([]*Item, error) {
	if !refresh && len(h.items) > 0 {
		return h.items, nil
	}

	prs, err := h.orchestrator.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	h.items = make([]*Item, 0, len(prs))
	for _, pr := range prs {
		item := NewItem(
			pr.Id(),
			pr.Title(),
			pr.Branch(),
			pr.State().String(),
			pr.LinkId(),
		)
		h.items = append(h.items, item)
	}

	return h.items, nil
}

func (h *chainAdaptor) GetItemsLinkedTo(item *Item, refresh bool) ([]*Item, error) {
	if !refresh && h.linkedItems[item.Id()] != nil {
		return h.linkedItems[item.Id()], nil
	}

	prs, err := h.orchestrator.GetPrsLinkedTo(item.Id())
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0, len(prs))
	for _, p := range prs {
		items = append(items, NewItem(
			p.Id(),
			p.Title(),
			p.Branch(),
			p.State().String(),
			p.LinkId(),
		))
	}

	h.linkedItems[item.Id()] = items
	return items, nil
}
