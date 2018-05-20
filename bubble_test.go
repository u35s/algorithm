package sort

import "testing"

func Test_Bubble(t *testing.T) {
	data := make([]int, len(arr))
	copy(data, arr)
	Bubble(data)
	t.Log(data)
}
