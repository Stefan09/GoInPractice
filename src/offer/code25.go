package offer

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//特殊情况判断
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//非去重合并
	newHeader := &ListNode{Val:-1, Next:nil}
	ptrWork := newHeader
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			ptrWork.Next = l1
			l1 = l1.Next
			ptrWork = ptrWork.Next
		}else{
			ptrWork.Next = l2
			l2 = l2.Next
			ptrWork = ptrWork.Next
		}
	}
	if l1 != nil {
		ptrWork.Next = l1
	}
	if l2 != nil {
		ptrWork.Next = l2
	}
	return newHeader.Next
}
