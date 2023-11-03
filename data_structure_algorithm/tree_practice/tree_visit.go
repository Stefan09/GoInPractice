package tree_practice

import (
	"container/list"
	"sort"
)

/**
题目描述
分别按照二叉树先序，中序和后序打印所有的节点。
*/
func threeOrders(root *TreeNode) [][]int {
	// write code here
	res := make([][]int, 3)
	res[0] = tlr(root)
	res[1] = ltr(root)
	res[2] = lrt(root)

	return res
}

//先序
func tlr(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, tlr(root.Left)...)
	res = append(res, tlr(root.Right)...)

	return res
}

//中序
func ltr(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, ltr(root.Left)...)
	res = append(res, root.Val)
	res = append(res, ltr(root.Right)...)

	return res
}

//后序
func lrt(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, lrt(root.Left)...)
	res = append(res, lrt(root.Right)...)
	res = append(res, root.Val)

	return res
}

/**
题目描述
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},

该二叉树层序遍历的结果是
[
[3],
[9,20],
[15,7]
]

示例1
输入
复制
{1,2}
返回值
[[1],[2]]

示例2
输入
{1,2,3,4,#,#,5}
返回值
复制
[[1],[2,3],[4,5]]
*/
type LevelNode struct {
	*TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	resMap := make(map[int][]int)
	res := make([][]int, 0)

	levelList := list.New()
	levelList.PushBack(&LevelNode{
		TreeNode: root,
		level:    0,
	})

	for levelList.Len() != 0 {
		curNode := levelList.Remove(levelList.Front()).(*LevelNode)
		resMap[curNode.level] = append(resMap[curNode.level], curNode.Val)
		if curNode.Left != nil {
			levelList.PushBack(&LevelNode{
				TreeNode: curNode.Left,
				level:    curNode.level + 1,
			})
		}
		if curNode.Right != nil {
			levelList.PushBack(&LevelNode{
				TreeNode: curNode.Right,
				level:    curNode.level + 1,
			})
		}
	}

	var keys []int
	for level := range resMap {
		keys = append(keys, level)
	}
	sort.Ints(keys)

	for key := range keys {
		res = append(res, resMap[key])
	}

	return res
}
