package leetcode

/*滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]

提示：

1 <= nums.length <= 10^5
-104 <= nums[i] <= 10^4
1 <= k <= nums.length
*/

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	maxList := make([]int, 0, len(nums)-k+1)
	queue := list.New()

	for i := 0; i < len(nums); i++ {
		// 删除队列中所有比当前元素小的元素
		for queue.Len() > 0 && nums[i] >= nums[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}
		// 将当前元素的索引加入队列
		queue.PushBack(i)

		// 删除队列中不在当前窗口范围内的元素
		if queue.Front().Value.(int) == i-k {
			queue.Remove(queue.Front())
		}

		// 当窗口形成后，将队列头部元素加入结果列表
		if i >= k-1 {
			maxList = append(maxList, nums[queue.Front().Value.(int)])
		}
	}
	return maxList
}
