package chain

import "testing"

func TestFindLink(t *testing.T) {
	body := "do not merge until #123 is released"
	want := uint(123)

	got := findLinkedPr(body)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
