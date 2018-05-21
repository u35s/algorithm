## golang map 支持key模糊搜索
go get github.com/u35s/searchmap
```
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/u35s/searchmap"
)

var keys []string
var m searchmap.SearchMap

func main() {
	m.Init()
	data, _ := ioutil.ReadFile("../data.txt")
	keys = strings.Split(string(data), ",")
	for i := range keys {
		m.Add(keys[i], keys[i])
	}
	fmt.Printf("数据大小:%v,索引大小:%v\n", m.Size(), m.IdxSize())
	fmt.Println("多的搜索结果:")
	fmt.Println(m.Search("多"))
}
```
