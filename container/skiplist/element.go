package skiplist

import (
	"math/rand"
	"time"
)

const SKIPLIST_MAXLEVEL = 32
const SKIPLIST_BRANCH = 4

type skiplistLevel struct {
	forward *Element
	span    int
}

type Element struct {
	Value    Interface
	backward *Element
	level    []*skiplistLevel
}

func (e *Element) Next() *Element {
	return e.level[0].forward
}

func (e *Element) Prev() *Element {
	return e.backward
}

func newElement(level int, v Interface) *Element {
	slLevels := make([]*skiplistLevel, level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skiplistLevel)
	}

	return &Element{
		Value:    v,
		backward: nil,
		level:    slLevels,
	}
}

func randomLevel() int {
	level := 1
	for (rand.Int31()&0xFFFF)%SKIPLIST_BRANCH == 0 {
		level += 1
	}

	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
