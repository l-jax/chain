package tui

import "chain/chain"

type chainAdaptor struct {
	orchestrator *chain.Orchestrator
	items        []*pr
	linkedItems  map[uint][]*pr
}

func newAdaptor() *chainAdaptor {
	return &chainAdaptor{
		orchestrator: chain.InitOrchestrator(),
		items:        []*pr{},
		linkedItems:  make(map[uint][]*pr),
	}
}

func (h *chainAdaptor) ListItems(refresh bool) ([]*pr, error) {
	if !refresh && len(h.items) > 0 {
		return h.items, nil
	}

	prs, err := h.orchestrator.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	h.items = make([]*pr, 0, len(prs))
	for _, pr := range prs {
		item := newPr(
			pr.Id(),
			pr.Branch(),
			pr.Title(),
			pr.Body(),
			pr.State().String(),
			pr.LinkId(),
		)
		h.items = append(h.items, item)
	}

	return h.items, nil
}

func (h *chainAdaptor) GetItemsLinkedTo(item *pr, refresh bool) ([]*pr, error) {
	if !refresh && h.linkedItems[item.Id()] != nil {
		return h.linkedItems[item.Id()], nil
	}

	prs, err := h.orchestrator.GetPrsLinkedTo(item.Id())
	if err != nil {
		return nil, err
	}

	items := make([]*pr, 0, len(prs))
	for _, p := range prs {
		items = append(items, newPr(
			p.Id(),
			p.Branch(),
			p.Title(),
			p.Body(),
			p.State().String(),
			p.LinkId(),
		))
	}

	h.linkedItems[item.Id()] = items
	return items, nil
}
