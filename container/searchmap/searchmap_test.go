package searchmap

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var keys []string
var m SearchMap

func init() {
	m.Init()
	data, _ := ioutil.ReadFile("data.txt")
	keys = strings.Split(string(data), ",")
	for i := range keys {
		m.Add(keys[i], keys[i])
	}
	fmt.Printf("数据大小:%v,索引大小:%v\n", m.Size(), m.IdxSize())
	res, ok := m.Search("多")
	fmt.Println("长度:", len(res), ok)
	fmt.Println(res)
}

func Benchmark_Search(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Search("艾西")
	}
}
