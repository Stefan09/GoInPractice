package offer

import "container/list"

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*/

//原地反转
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var curr, newHeader *ListNode = head, nil
	for curr != nil {
		next := curr.Next
		curr.Next = newHeader
		newHeader = curr
		curr = next
		/* 化简写法，可读性不佳
		curr, curr.Next, newHeader = curr.Next, newHeader, curr
		*/
	}
	return newHeader
}

//借助栈
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	stack := list.New()
	iter := head
	flag := true
	for iter != nil {
		stack.PushFront(iter)
		iter = iter.Next
	}
	var newHeader, ptrTail *ListNode
	for stack.Len() != 0 {
		iter = stack.Remove(stack.Front()).(*ListNode)
		if flag {
			newHeader = iter
			ptrTail = iter
			flag = false
		}
		ptrTail.Next = iter
		ptrTail = ptrTail.Next
	}
	ptrTail.Next = nil
	return newHeader
}
