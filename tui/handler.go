package tui

import (
	"chain/chain"
	"fmt"
)

type handler struct {
	links  []pr
	chains map[uint][]pr
}

func initHandler() *handler {
	return &handler{
		links:  []pr{},
		chains: make(map[uint][]pr),
	}
}

func (h *handler) FetchOpen(refresh bool) ([]pr, error) {
	if !refresh && len(h.links) > 0 {
		return h.links, nil
	}

	chainHandler := chain.NewChainHandler()
	pulls, err := chainHandler.GetPullRequests()
	if err != nil {
		return nil, err
	}

	h.links = make([]pr, len(pulls))
	for i, pull := range pulls {
		link, err := mapPr(pull)
		if err != nil {
			return nil, err
		}
		h.links[i] = *link
	}
	return h.links, nil
}

func (h *handler) FetchChain(link pr, refresh bool) ([]pr, error) {
	if !refresh && h.chains[link.id] != nil {
		return h.chains[link.id], nil
	}

	chainHandler := chain.NewChainHandler()
	pull, err := chainHandler.GetChain(link.Id())
	if err != nil {
		return nil, err
	}

	links := make([]pr, 0, len(pull))
	for _, pr := range pull {
		link, err := mapPr(pr)
		if err != nil {
			return nil, err
		}
		links = append(links, *link)
	}

	h.chains[link.id] = links
	return links, nil
}

func mapPr(pr *chain.PullRequest) (*pr, error) {
	label, err := mapLabel(pr.State())
	if err != nil {
		return nil, err
	}
	link := InitPr(
		pr.Title(),
		pr.Body(),
		pr.Branch(),
		pr.Number(),
		pr.Chain(),
		label,
	)
	return &link, nil
}

func mapLabel(state chain.State) (state, error) {
	switch state {
	case chain.StateOpen:
		return open, nil
	case chain.StateBlocked:
		return blocked, nil
	case chain.StateMerged:
		return merged, nil
	case chain.StateReleased:
		return released, nil
	case chain.StateClosed:
		return closed, nil
	default:
		return open, fmt.Errorf("unknown state: %s", state)
	}
}
