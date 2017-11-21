package sort

import "testing"

var arr = []int{2, 1, 0, 3, 5, 4}

func Test_Insertion(t *testing.T) {
	data := make([]int, len(arr))
	copy(data, arr)
	Insertion(data)
	t.Log(data)
}

func Test_Bubble(t *testing.T) {
	data := make([]int, len(arr))
	copy(data, arr)
	Bubble(data)
	t.Log(data)
}
