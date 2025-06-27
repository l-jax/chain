package tui

type listMsg struct {
	items []*Item
}

type detailMsg struct {
	item *Item
}

type chainMsg struct {
	items []*Item
}

type errMsg struct {
	err error
}
