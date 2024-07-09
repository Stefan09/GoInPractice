package leetcode

/**
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]


链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	// 设置一个虚拟头
	fakeHead := &ListNode{Val: -1, Next: head}
	pre, slow, fast := fakeHead, head, head
	i := 1
	for ; i <= n-1 && fast != nil; i++ {
		fast = fast.Next
	}
	if i < n { // n大于链表长度
		return head
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		pre = pre.Next
	}
	pre.Next = slow.Next

	return fakeHead.Next
}
