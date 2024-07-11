package leetcode

import "container/list"

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5


提示：
树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q, depth := list.New(), 1
	q.PushBack(root)
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			cur := q.Remove(q.Front()).(*TreeNode)
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		depth++
	}
	return depth
}
