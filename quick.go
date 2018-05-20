package sort

// 快速排序
func Quick(arr []int) {
	if len(arr) < 2 {
		return
	}
	//tmp := make([]int, len(arr))
	//copy(tmp, arr)
	pivot := partition(arr)
	//fmt.Println(tmp, pivot, arr)
	if pivot > 0 {
		Quick(arr[:pivot])
	}
	if pivot+1 < len(arr) {
		Quick(arr[pivot+1:])
	}
}

// 返回枢轴pivot
func partition(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	// 枢轴默认取第一个元素

	// 最好的情况下,每次划分的很均匀,递归树的深度为log2n,
	// 时间复杂度为O(nlogn)

	// 最坏情况为待排序的序列为正序或逆序,每次划分只得到一个比上一次划分少一个记录的子序列
	// 这种情况下要执行n-1个递归调用,且第i次需要经过n-1次比较才能找到枢纽的位置,
	// 因此比较次数为(n-1)+(n-2)+1=n(n-1)/2,最终其时间复杂度为O(n^2)
	pivotKey := arr[0]
	low, high := 0, len(arr)-1

	for low < high {
		for low < high && arr[high] >= pivotKey {
			high--
		}
		// 将比枢轴小的交换到低端
		arr[low], arr[high] = arr[high], arr[low]

		for low < high && arr[low] <= pivotKey {
			low++
		}
		// 将比枢轴大的交换到高端
		arr[low], arr[high] = arr[high], arr[low]
	}
	return low
}
