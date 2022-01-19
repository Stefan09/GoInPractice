package search_sort

/**
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。

请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


提示：

1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4
*/

//二分法变式，需要确认mid&target落在哪个片段，从而确认二分之真如何移动
func search(nums []int, target int) int {
	//异常判断
	if len(nums) == 0 {
		return -1
	}

	//二分变式
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[len(nums)-1] { //前半段
			if target > nums[len(nums)-1] && target > nums[mid] {

			} else {

			}
		} else { //后半段
			if target < nums[len(nums)-1] && target > nums[mid] {

			} else {

			}
		}
	}
	return 0
}
