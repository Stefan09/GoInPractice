package sort_practice

func QuickSort(arr []int, low, high int) {
	if len(arr) == 0 || low >= high {
		return
	}

	partitionIndex := partition(arr, low, high)
	QuickSort(arr, low, partitionIndex-1)
	QuickSort(arr, partitionIndex+1, high)
}

func partition(arr []int, low, high int) int {
	if low == high {
		return low
	}

	temp := arr[low] //设置基准点（基准点的选取可以使用多种方法增加随机性）
	for low < high {
		for low < high && arr[high] > temp {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] < temp {
			low++
		}
		arr[high] = arr[low]
	}

	arr[low] = temp
	return low
}
