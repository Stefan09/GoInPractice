package search_sort_practice

/**
题目描述
有一个整数数组，请你根据快速排序的思路，找出数组中第K大的数。

给定一个整数数组a,同时给定它的大小n和要找的K(K在1到n之间)，请返回第K大的数，保证答案存在。

示例1
输入
[1,3,5,2,2],5,3
返回值
2
*/

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func FindKth(a []int, n int, K int) int {
	// 利用快排partition，每次确定一个元素位置且左右可区分大小
	if len(a) == 0 || n == 0 || K > n {
		return -1
	}

	left, right := 0, n-1
	temp := a[left]
	for left < right {
		for left < right && a[right] <= temp {
			right--
		}
		a[left] = a[right]

		for left < right && a[left] > temp {
			left++
		}
		a[right] = a[left]
	}
	a[left] = temp

	if left == K {
		return a[left]
	} else if left > K {
		return FindKth(a[:left], left, K)
	} else {
		return FindKth(a[left+1:], n-left-1, K)
	}

}
