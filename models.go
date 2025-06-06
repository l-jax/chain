package main

import (
	"slices"
)

type PullRequest struct {
	title      string
	branch     string
	conditions []Condition
}

type Condition interface {
	success() bool
}

type StateCondition struct {
	currentState string
	targetState  string
}

func (s *StateCondition) success() bool {
	return s.currentState == s.targetState
}

type DependencyCondition struct {
	dependency *PullRequest
	conditions []Condition
}

func (d *DependencyCondition) success() bool {
	for _, v := range slices.All(d.conditions) {
		if !v.success() {
			return false
		}
	}
	return true
}
