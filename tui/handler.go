package tui

import (
	"chain/chain"
)

type handler struct {
	links  []chain.Pr
	chains map[uint][]chain.Pr
}

func initHandler() *handler {
	return &handler{
		links:  []chain.Pr{},
		chains: make(map[uint][]chain.Pr),
	}
}

func (h *handler) FetchOpen(refresh bool) ([]chain.Pr, error) {
	if !refresh && len(h.links) > 0 {
		return h.links, nil
	}

	chainHandler := chain.InitOrchestrator()
	pulls, err := chainHandler.ListOpenPrs()
	if err != nil {
		return nil, err
	}

	h.links = make([]chain.Pr, len(pulls))
	for i, pull := range pulls {
		h.links[i] = *pull
	}
	return h.links, nil
}

func (h *handler) FetchChain(link chain.Pr, refresh bool) ([]chain.Pr, error) {
	if !refresh && h.chains[link.Id()] != nil {
		return h.chains[link.Id()], nil
	}

	chainHandler := chain.InitOrchestrator()
	pull, err := chainHandler.GetPrsLinkedTo(link.Id())
	if err != nil {
		return nil, err
	}

	links := make([]chain.Pr, 0, len(pull))
	for _, pr := range pull {
		links = append(links, *pr)
	}

	h.chains[link.Id()] = links
	return links, nil
}
