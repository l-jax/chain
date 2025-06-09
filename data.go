package main

import "github.com/charmbracelet/bubbles/list"

func (c *Chain) initLists() {
	first := NewPull("merge this first", "my-first-branch", StateMerged, nil)
	second := NewPull("merge this second", "my-second-branch", StateOpen, &first)
	third := NewPull("merge this third", "my-third-branch", StateOpen, &second)

	c.groups = []group{
		newGroup(StateOpen),
		newGroup(StateMerged),
		newGroup(StateReleased),
	}

	c.groups[StateOpen].list.Title = StateOpen.String()
	c.groups[StateOpen].list.SetItems([]list.Item{
		second,
		third,
	})

	c.groups[StateMerged].list.Title = StateMerged.String()
	c.groups[StateMerged].list.SetItems([]list.Item{
		first,
	})
	
	c.loaded = true
}
