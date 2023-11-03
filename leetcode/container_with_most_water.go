package leetcode

/**
https://leetcode-cn.com/problems/container-with-most-water/
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例 2：
输入：height = [1,1]
输出：1

提示：
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

// brute force
/*func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			a := (j - i) * min(height[i], height[j])
			if a > max {
				max = a
			}
		}
	}
	return max
}*/

// double pointer
func maxArea(height []int) int {
	max := 0
	for left, right := 0, len(height)-1; left < right; {
		a := (right -left) * min(height[left], height[right])
		if a > max {
			max = a
		}
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return max
}
