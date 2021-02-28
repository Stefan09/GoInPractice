package offer

import "container/list"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

//数组翻转
func ReversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	iter := head
	var res []int
	for iter != nil {
		res = append(res, iter.Val)
		iter = iter.Next
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

//stack 栈相当于头插法链表
func ReversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	iter := head
	stack := list.New()
	var res []int
	for iter != nil {
		stack.PushFront(iter.Val)
		iter = iter.Next
	}
	for stack.Len() != 0 {
		res = append(res, stack.Remove(stack.Front()).(int))
	}
	return res
}
