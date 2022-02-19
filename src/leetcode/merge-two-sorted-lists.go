package leetcode

/**
https://leetcode-cn.com/problems/merge-two-sorted-lists/

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{}
	t := h
	l1, l2 := list1, list2

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			t.Next = l1
			t = t.Next
			l1 = l1.Next
			t.Next = nil
		} else {
			t.Next = l2
			t = t.Next
			l2 = l2.Next
			t.Next = nil
		}
	}

	if l1 != nil {
		t.Next = l1
	}
	if l2 != nil {
		t.Next = l2
	}

	return h.Next
}
