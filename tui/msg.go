package tui

type listMsg struct {
	items []*pr
}

type detailMsg struct {
	item *pr
}

type chainMsg struct {
	items []*pr
}

type errMsg struct {
	err error
}
