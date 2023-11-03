package search_sort

/**
冒泡排序的理解方式很多，例如：
1. 大数下沉
2. 小数上浮
*/
func BubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
