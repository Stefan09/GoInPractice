package offer

import "sort"

/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
*/

//1. 利用递增特性 二分法
//2. 边界
func MinArray(numbers []int) int {
	//Left为数组最左端下标，Right为数组最右端下标
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if numbers[left] < numbers[right] {
			return numbers[left]
		}
		if numbers[mid] > numbers[left] {
			left = mid + 1
		} else if numbers[mid] < numbers[left] {
			right = mid
		} else {
			left++
		}
	}
	return numbers[left]
}

func MinArray1(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
