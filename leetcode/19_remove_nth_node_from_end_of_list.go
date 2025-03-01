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
1 <= n <= sz 不会有n的非法输入
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fakeHead := &ListNode{Val: -1, Next: head}
	pre, cur, fast := fakeHead, head, head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		pre = pre.Next
		cur = cur.Next
		fast = fast.Next
	}
	pre.Next = cur.Next

	return fakeHead.Next
}
