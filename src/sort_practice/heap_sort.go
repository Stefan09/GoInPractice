package sort_practice

//调整：从上至下调整
//注意参数传递，需要限制下界
func HeapSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	buildHeap(arr)
	。
	//todo 根据需要使用堆
}

//建堆
func buildHeap(arr []int) {
	for iter := (len(arr) - 1) / 2; iter >= 0; iter-- {
		maxHeap(arr, iter, len(arr)-1)
	}
}

//堆调整
func maxHeap(arr []int, curRoot, end int) {
	left := 2*curRoot + 1
	right := 2*curRoot + 2
	max := curRoot

	if left < end && arr[left] > arr[max] {
		max = left
	}
	if right < end && arr[right] > arr[max] {
		max = right
	}

	if max != curRoot {
		arr[curRoot], arr[max] = arr[max], arr[curRoot]
		maxHeap(arr, max, end)
	}
}
