package tui

import (
	"chain/chain"
)

//TODO: map chain.Pr to tui model that implements list.Item

type chainAdaptor struct {
	orchestrator *chain.Orchestrator
	openPrs      []*chain.Pr
	linkedPrs    map[uint][]chain.Pr
}

func initChainAdaptor() *chainAdaptor {
	return &chainAdaptor{
		orchestrator: chain.InitOrchestrator(),
		openPrs:      []*chain.Pr{},
		linkedPrs:    make(map[uint][]chain.Pr),
	}
}

func (h *chainAdaptor) ListOpenPrs(refresh bool) ([]*chain.Pr, error) {
	if !refresh && len(h.openPrs) > 0 {
		return h.openPrs, nil
	}

	var err error
	h.openPrs, err = h.orchestrator.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	return h.openPrs, nil
}

func (h *chainAdaptor) GetPrsLinkedTo(pr *chain.Pr, refresh bool) ([]chain.Pr, error) {
	if !refresh && h.linkedPrs[pr.Id()] != nil {
		return h.linkedPrs[pr.Id()], nil
	}

	prs, err := h.orchestrator.GetPrsLinkedTo(pr.Id())
	if err != nil {
		return nil, err
	}

	linkedPrs := make([]chain.Pr, 0, len(prs))
	for _, p := range prs {
		linkedPrs = append(linkedPrs, *p)
	}

	h.linkedPrs[pr.Id()] = linkedPrs
	return linkedPrs, nil
}
