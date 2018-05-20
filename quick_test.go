package sort

import "testing"

var quick_arr1 = []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
var quick_arr2 = []int{2, 1, 0, 3, 5, 4, 33, 12, 49, 66, 204, 55, 43, 78}

func Test_Quick(t *testing.T) {
	data := make([]int, len(quick_arr1))
	copy(data, quick_arr1)
	Quick(data)
	t.Log(data)

	data = make([]int, len(quick_arr2))
	copy(data, quick_arr2)
	Quick(data)
	t.Log(data)
}
