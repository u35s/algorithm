package sort

import "testing"

func Test_Heap(t *testing.T) {
	data := make([]int, len(arr))
	copy(data, arr)
	Heap(data)
	t.Log(data)
}
