package sort

import "fmt"

// 堆排序
func Heap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		down(arr, i, n)
	}
}
func down(arr []int, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && arr[j2] < arr[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		//if !h.Less(j, i) {
		if arr[j] >= arr[i] {
			break
		}
		//h.Swap(i, j)
		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}
	return i > i0
}
