package offer

/**
题目描述
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值

[要求]
set和get方法的时间复杂度为O(1)
某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案

示例1
输入
复制
[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
返回值
复制
[1,-1]
说明
第一次操作后：最常使用的记录为("1", 1)
第二次操作后：最常使用的记录为("2", 2)，("1", 1)变为最不常用的
第三次操作后：最常使用的记录为("3", 2)，("1", 1)还是最不常用的
第四次操作后：最常用的记录为("1", 1)，("2", 2)变为最不常用的
第五次操作后：大小超过了3，所以移除此时最不常使用的记录("2", 2)，加入记录("4", 4)，并且为最常使用的记录，然后("3", 2)变为最不常使用的记录
*/

/**
解题思路：
1. get&set O(1)操作 ==> 利用hash结构存储（map）
2. lru最近最久未使用，存储内部元素“有序”
*/
func LRU(operators [][]int, k int) []int {
	res := make([]int, 0, len(operators))
	keys := make([]int, 0, k)
	values := make([]int, 0, k)
	for _, arr := range operators {
		if arr[0] == 1 {
			if len(keys) == k {
				keys = keys[1:]
				values = values[1:]
			}
			keys = append(keys, arr[1])
			values = append(values, arr[2])
		} else if arr[0] == 2 {
			var i int
			for i = 0; i < k; i++ {
				if i >= len(keys) {
					break
				}
				if arr[1] == keys[i] {
					break
				}
			}
			if i == k || i >= len(keys) {
				res = append(res, -1)
			} else {
				res = append(res, values[i])
				if i+1 < k {
					keys = append(keys[0:i], append(keys[i+1:], keys[i])...)
					values = append(values[0:i], append(values[i+1:], values[i])...)
				}
			}
		}
	}
	return res
}
