skiplist
===============

reference from gansidui [skiplist](https://github.com/gansidui/skiplist)

new feature
==============

* 允许重复值
* 结构简化
* 打印跳表结构
* 随机函数优化


Usage
===============

~~~Go

package main

import (
	"fmt"

	"github.com/gansidui/skiplist"
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

/* output:
length:51,45rank6:6
h-------------------------------------------------------------------------------------------------------------------------------------------------02----00
h-------47-46----------43-------40----------36-------------------29----27-------------------------18-------15-------------10----------------------02----00
h-49-48-47-46-45-45-44-43-42-41-40-39-38-37-36-35-34-33-32-31-30-29-28-27-26-25-24-23-22-21-20-19-18-17-16-15-14-13-12-11-10-09-08-07-06-05-04-03-02-01-00
length:50,44rank6:6
h----------------------------------------------------------------------------------------------------------------------------------------------02----00
h-------47-46-------43-------40----------36-------------------29----27-------------------------18-------15-------------10----------------------02----00
h-49-48-47-46-45-44-43-42-41-40-39-38-37-36-35-34-33-32-31-30-29-28-27-26-25-24-23-22-21-20-19-18-17-16-15-14-13-12-11-10-09-08-07-06-05-04-03-02-01-00
*/

~~~


License
===============

MIT
