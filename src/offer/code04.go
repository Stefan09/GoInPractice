package offer

/**
二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	//边界判断
	if len(matrix) == 0 {
		return false
	}

	var xIndex, yIndex = 0, len(matrix) - 1

	for xIndex < len(matrix[0]) && yIndex >= 0 {
		if target == matrix[yIndex][xIndex] {
			return true
		} else if target > matrix[yIndex][xIndex] {
			xIndex++
		} else {
			yIndex--
		}
	}
	return false
}
