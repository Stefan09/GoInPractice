package leetcode

import (
	"container/heap"
)

/**
https://leetcode-cn.com/problems/merge-k-sorted-lists/
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

// 参考两个链表merge，关键点在于如何获得多个链表中最小的结点
// 利用小顶堆/优先队列，弹出最小并且将后继结点加入堆中
func mergeKLists(lists []*ListNode) *ListNode {
	res := new(ListNode)
	p := res

	h := new(ListNodeMinHeap)
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		p.Next = node
		p = p.Next
	}

	return res.Next
}
