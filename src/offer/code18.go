package offer

/**
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
*/

func deleteNode(head *ListNode, val int) *ListNode {
	//边界情况
	if head == nil {
		return nil
	}
	//使用头副本，便于返回链表头
	copyHead := &ListNode{Next: head,}
	pre := copyHead
	iter := copyHead.Next
	for iter != nil {
		if iter.Val != val {
			pre = iter
			iter = iter.Next
			continue
		}
		pre.Next = iter.Next
		iter = iter.Next
	}
	return copyHead.Next
}
