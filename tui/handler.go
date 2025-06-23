package tui

import (
	"chain/chain"
	"fmt"
)

type handler struct {
	links  []Link
	chains map[uint][]Link
}

func initHandler() *handler {
	return &handler{
		links:  []Link{},
		chains: make(map[uint][]Link),
	}
}

func (h *handler) FetchOpen(refresh bool) ([]Link, error) {
	if !refresh && len(h.links) > 0 {
		return h.links, nil
	}

	chainHandler := chain.NewChainHandler()
	pulls, err := chainHandler.GetPullRequests()
	if err != nil {
		return nil, err
	}

	h.links = make([]Link, len(pulls))
	for i, pull := range pulls {
		link, err := mapPr(pull)
		if err != nil {
			return nil, err
		}
		h.links[i] = *link
	}
	return h.links, nil
}

func (h *handler) FetchChain(link Link, refresh bool) ([]Link, error) {
	if !refresh && h.chains[link.id] != nil {
		return h.chains[link.id], nil
	}

	chainHandler := chain.NewChainHandler()
	pull, err := chainHandler.GetChain(link.Id())
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0, len(pull))
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

func mapPr(pr *chain.PullRequest) (*Link, error) {
	label, err := mapLabel(pr.State())
	if err != nil {
		return nil, err
	}
	link := NewLink(
		pr.Title(),
		pr.Body(),
		pr.Branch(),
		pr.Number(),
		pr.Chain(),
		label,
	)
	return &link, nil
}

func mapLabel(state chain.State) (label, error) {
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
