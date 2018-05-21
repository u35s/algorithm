package main

import (
	"fmt"

	"github.com/u35s/algorithm/container/skiplist"
)

type User struct {
	score int
}

func (u *User) String() string {
	if u.score < 10 {
		return fmt.Sprintf("-0%v", u.score)
	}
	return fmt.Sprintf("-%v", u.score)
}

func (u *User) Less(other interface{}) bool {
	return u.score >= other.(*User).score
}

func main() {
	us := make([]*User, 50)
	for i := 0; i < 50; i++ {
		us[i] = &User{i}
	}
	sl := skiplist.New()
	for i := 0; i < len(us); i++ {
		sl.Insert(us[i])
	}
	s := &User{45}
	sl.Insert(s)
	e := sl.GetElementByRank(6)
	fmt.Printf("length:%v,45rank6:%v\n", sl.Len(), sl.GetRank(e.Value))
	fmt.Printf(sl.String())
	sl.Remove(e)
	e = sl.GetElementByRank(6)
	fmt.Printf("length:%v,44rank6:%v\n", sl.Len(), sl.GetRank(e.Value))
	fmt.Printf(sl.String())
}
