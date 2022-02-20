package leetcode

/**
 * This file is used to define data structure and useful method
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type char uint8

// list node min heap
type ListNodeMinHeap []*ListNode

func (m ListNodeMinHeap) Len() int {
	return len(m)
}

func (m ListNodeMinHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m ListNodeMinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *ListNodeMinHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *ListNodeMinHeap) Pop() interface{} {
	old := *m
	e := old[len(old)-1]
	*m = old[:len(old)-1]

	return e
}

/************** useful method ****************/
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
