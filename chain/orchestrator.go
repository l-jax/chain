package chain

import (
	"regexp"
	"strconv"
)

type orchestrator struct {
	adaptor adaptor
}

func NewOrchestrator(adaptor adaptor) *orchestrator {
	return &orchestrator{
		adaptor: adaptor,
	}
}

func (o *orchestrator) GetPullRequests() ([]*Pull, error) {
	pulls, err := o.adaptor.listPullRequests()

	if err != nil {
		return nil, err
	}

	return pulls, nil
}

func (o *orchestrator) GetChain(number uint) ([]*Pull, error) {
	pull, err := o.adaptor.getPullRequest(number)

	chain := make([]*Pull, 0)
	chain = append(chain, pull)

	if err != nil {
		return nil, err
	}

	linkedPull, err := o.findLinkedPull(pull)

	if err != nil {
		return nil, err
	}

	if linkedPull == nil {
		return chain, nil
	}

	chain = append(chain, linkedPull)

	return chain, nil
}

func (o *orchestrator) findLinkedPull(pull *Pull) (*Pull, error) {
	re := regexp.MustCompile(`do not merge until #(\d+)`)
	match := re.FindStringSubmatch(pull.Body())
	if match == nil {
		return nil, nil
	}

	linkedPullNumber := match[1]
	linkedPullNumberUint, err := strconv.ParseUint(linkedPullNumber, 10, 32)
	if err != nil {
		return nil, err
	}
	return o.adaptor.getPullRequest(uint(linkedPullNumberUint))
}

func GetOpen() []Pull {
	// TODO: open logic
	return []Pull{
		{"my pull request", "some-branch", "some body", StateOpen, 1, nil},
		{"code", "some-other-branch", "some body", StateOpen, 2, nil},
	}
}

func GetChain(branch string) []Pull {
	// TODO: chain logic
	three := Pull{"remove something", "some-branch-123", "some body", StateReleased, 1, nil}
	two := Pull{"add something", "my-branch", "some body", StateMerged, 2, &three}
	one := Pull{"do something", branch, "some body", StateOpen, 3, &two}
	return []Pull{
		one,
		two,
		three,
	}
}
