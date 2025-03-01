package leetcode

/*移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
提示:
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
进阶：你能尽量减少完成的操作次数吗？
*/

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

/* 写得不漂亮 但是很奇怪效率却挺高
func moveZeroes(nums []int) {
	zeroPtr, nonZeroPtr := 0, 0
	for zeroPtr < len(nums) {
		if nums[zeroPtr] != 0 {
			zeroPtr++
			nonZeroPtr++
		} else {
			for nonZeroPtr < len(nums) {
				if nums[nonZeroPtr] != 0 {
					nums[zeroPtr], nums[nonZeroPtr] = nums[nonZeroPtr], nums[zeroPtr]
					break
				}
				nonZeroPtr++
			}
			zeroPtr++
		}
	}
}
*/
