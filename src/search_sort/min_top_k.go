package search_sort

import "sort"

/**
题目描述
输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。
示例1
输入
复制
[4,5,1,6,2,7,3,8],4
返回值
复制
[1,2,3,4]
*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	//build big-heap
	if k == 0 || len(input) == 0 || k > len(input) {
		return nil
	}

	heap := make([]int, k)
	for i := 0; i < k; i++ {
		heap[i] = input[i]
	}
	for i := (k - 1) / 2; i >= 0; i-- {
		adjust(heap, i)
	}

	//iter input & adjust
	for i := k; i < len(input); i++ {
		if input[i] < heap[0] {
			heap[0] = input[i]
			adjust(heap, 0)
		}
	}

	sort.Ints(heap)
	return heap
}

func adjust(heap []int, index int) {
	max := index
	left, right := 2*index+1, 2*index+2
	if left < len(heap) && heap[left] > heap[max] {
		max = left
	}
	if right < len(heap) && heap[right] > heap[max] {
		max = right
	}
	if max != index {
		heap[max], heap[index] = heap[index], heap[max]
		index = max
		adjust(heap, index)
	}
}
