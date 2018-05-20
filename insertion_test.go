package sort

import "testing"

func Test_Insertion(t *testing.T) {
	data := make([]int, len(arr))
	copy(data, arr)
	Insertion(data)
	t.Log(data)
}
