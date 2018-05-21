package skiplist

import (
	"bytes"
)

type Interface interface {
	Less(other interface{}) bool
	String() string
}

type SkipList struct {
	header *Element
	tail   *Element
	length int
	level  int
}

func New() *SkipList {
	return &SkipList{
		header: newElement(SKIPLIST_MAXLEVEL, nil),
		tail:   nil,
		length: 0,
		level:  1,
	}
}

func (sl *SkipList) Init() *SkipList {
	sl.header = newElement(SKIPLIST_MAXLEVEL, nil)
	sl.tail = nil
	sl.length = 0
	sl.level = 1
	return sl
}

func (sl *SkipList) Front() *Element {
	return sl.header.level[0].forward
}

func (sl *SkipList) Back() *Element {
	return sl.tail
}

func (sl *SkipList) Len() int {
	return sl.length
}

func (sl *SkipList) Insert(v Interface) *Element {
	x := sl.header
	update := make([]*Element, SKIPLIST_MAXLEVEL)
	rank := make([]int, SKIPLIST_MAXLEVEL)
	for i := sl.level - 1; i >= 0; i-- {
		if i == sl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	level := randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.header
			update[i].level[i].span = sl.length
		}
		sl.level = level
	}

	x = newElement(level, v)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		x.level[i].span = update[i].level[i].span - rank[0] + rank[i]
		update[i].level[i].span = rank[0] - rank[i] + 1
	}

	for i := level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == sl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		sl.tail = x
	}
	sl.length++

	return x
}

func (sl *SkipList) Remove(e *Element) {
	x := sl.header
	update := make([]*Element, SKIPLIST_MAXLEVEL)
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.Value.Less(e.Value) {
			if x.level[i].forward == e {
				break
			}
			x = x.level[i].forward
		}
		update[i] = x
	}
	for i := sl.level - 1; i >= 0; i-- {
		if update[i].level[i].forward == e {
			update[i].level[i].span += e.level[i].span - 1
			update[i].level[i].forward = e.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if e.level[0].forward != nil {
		e.level[0].forward.backward = e.backward
	} else {
		sl.tail = e.backward
	}

	for sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList) Find(v Interface) *Element {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
			if x.level[i].forward.Value == v {
				return x
			}
			x = x.level[i].forward
		}
	}
	return nil
}

func (sl *SkipList) GetRank(v Interface) int {
	x := sl.header
	rank := 0
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
			rank += x.level[i].span
			if x.level[i].forward.Value == v {
				return rank
			}
			x = x.level[i].forward
		}
	}
	return 0
}

func (sl *SkipList) GetElementByRank(rank int) *Element {
	x := sl.header
	traversed := 0
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && traversed+x.level[i].span <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}

	return nil
}

func (sl *SkipList) String() string {
	buf := bytes.Buffer{}
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		buf.WriteString("h")
		for x.level[i].forward != nil {
			for j := 1; j < x.level[i].span; j++ {
				buf.WriteString("---")
			}
			x = x.level[i].forward
			buf.WriteString(x.Value.String())
			buf.WriteString("")
		}
		buf.WriteString("\n")
		x = sl.header
	}
	return buf.String()
}
