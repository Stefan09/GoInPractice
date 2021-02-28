package linklist_practice

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, quick := head, head
	for slow != nil && quick != nil {
		slow = slow.Next
		// 连续前进两步前需要合法性判断
		if quick.Next == nil {
			return false
		}
		quick = quick.Next.Next
		if slow == quick {
			return true
		}
	}

	return false
}
