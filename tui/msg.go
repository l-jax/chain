package tui

type listMsg struct {
	items []*Item
}

type detailMsg struct {
	item *Item
}

type tableMsg struct {
	items []*Item
}

type errMsg struct {
	err error
}
