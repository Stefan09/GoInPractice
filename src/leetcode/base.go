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

/************** useful method ****************/
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
