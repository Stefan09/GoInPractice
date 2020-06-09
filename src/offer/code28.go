package offer

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

*/

/**
method1: 递归
注意各种边界以及函数意义
*/
func isSymmetric(root *TreeNode) bool {
	//空节点
	if root == nil {
		return true
	}
	return judgeSymmetric(root.Left, root.Right)
}

func judgeSymmetric(left *TreeNode, right *TreeNode) bool {
	//边界条件
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil || left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return judgeSymmetric(left.Right, right.Left) && judgeSymmetric(left.Left, right.Right)
}

/**
method2: 中序遍历, 双指针检查
*/
var visitList []int

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	inOrderVisit(root)
	left, right := 0, len(visitList)-1
	for left < right {
		if visitList[left] != visitList[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func inOrderVisit(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderVisit(root.Left)
	visitList = append(visitList, root.Val)
	inOrderVisit(root.Right)
}
