package tui

type resetMsg struct{}

type listMsg struct {
	items []*Item
}

type detailMsg struct {
	item *Item
}

type tableLoadMsg struct{}

type tableMsg struct {
	items []*Item
}

type errMsg struct {
	err error
}
