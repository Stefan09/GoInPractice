package linklist

import "offer"

/**
题目描述
输入一个链表，反转链表后，输出新链表的表头。

示例1
输入
复制
{1,2,3}
返回值
复制
{3,2,1}
*/

// 头插法
func ReverseListHead(pHead *offer.ListNode) *offer.ListNode {
	if pHead == nil {
		return nil
	}

	// create new head
	newHead := &offer.ListNode{
		Val:  0,
		Next: nil,
	}

	// visit & rebuild
	iter := pHead
	for iter != nil {
		temp := iter.Next
		iter.Next = newHead.Next
		newHead.Next = iter
		iter = temp
	}

	return newHead.Next
}

// 尾插法
func ReverseListTail(pHead *offer.ListNode) *offer.ListNode {
	// 异常判断
	if pHead == nil {
		return nil
	}

	// 头结点单独处理
	newHead, iter := pHead, pHead.Next
	newHead.Next = nil

	// 遍历反转
	// 断链 重构 指针移动
	for iter != nil {
		temp := iter.Next
		iter.Next = newHead
		newHead = iter
		iter = temp
	}

	// 返回新头结点
	return newHead
}
