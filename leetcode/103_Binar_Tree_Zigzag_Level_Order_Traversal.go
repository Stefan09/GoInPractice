package leetcode

import "container/list"

/*
二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	leftToRight := 1

	// 队列中节点顺序和bfs一致，只是读取顺序做调整
	for queue.Len() != 0 {
		length := queue.Len()
		level := make([]int, 0)
		for i := 0; i < length; i++ {
			if leftToRight == 1 {
				tn := queue.Remove(queue.Front()).(*TreeNode)
				level = append(level, tn.Val)
				if tn.Left != nil {
					queue.PushBack(tn.Left)
				}
				if tn.Right != nil {
					queue.PushBack(tn.Right)
				}
			} else {
				tn := queue.Remove(queue.Back()).(*TreeNode)
				level = append(level, tn.Val)
				if tn.Right != nil {
					queue.PushFront(tn.Right)
				}
				if tn.Left != nil {
					queue.PushFront(tn.Left)
				}
			}
		}
		leftToRight ^= 1
		res = append(res, level)
	}
	return res
}
