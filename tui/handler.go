package tui

import "chain/chain"

type handler struct {
	links []Link
}

func initHandler() *handler {
	return &handler{
		links: []Link{},
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
		h.links[i] = NewLink(
			pull.Title(),
			pull.Body(),
			pull.Branch(),
			pull.Number(),
			pull.Chain(),
			label(pull.State()),
		)
	}
	return h.links, nil
}

func (h *handler) FetchChain(link Link) ([]Link, error) {
	chainHandler := chain.NewChainHandler()
	pull, err := chainHandler.GetChain(link.Id())
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0, len(pull))
	for _, pr := range pull {
		links = append(links, NewLink(
			pr.Title(),
			pr.Body(),
			pr.Branch(),
			pr.Number(),
			pr.Chain(),
			label(pr.State()),
		))
	}
	return links, nil
}
