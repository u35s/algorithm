package sort

// 冒泡
func Bubble(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n-1; i++ {
		change := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				change = true
			}
		}
		if !change {
			break
		}
	}
}
