package search_sort_practice

// 1. 划分子数组
// 2. 子数组排序
// 3. merge
func MergeSort(arr []int) []int {
	// 注意递归出口
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	arrLeft, arrRight := arr[:mid], arr[mid:]
	left := MergeSort(arrLeft)
	right := MergeSort(arrRight)
	result := merge(left, right)

	return result
}

// 子数组归并
func merge(left, right []int) []int {
	result := make([]int, 0)
	leftIter, rightIter := 0, 0

	for leftIter < len(left) && rightIter < len(right) {
		if left[leftIter] <= right[rightIter] {
			result = append(result, left[leftIter])
			leftIter++
		} else {
			result = append(result, right[rightIter])
			rightIter++
		}
	}

	if leftIter < len(left) {
		result = append(result, left[leftIter:]...)
	}

	if rightIter < len(right) {
		result = append(result, right[rightIter:]...)
	}

	return result
}
