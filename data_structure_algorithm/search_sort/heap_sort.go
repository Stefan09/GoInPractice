/** 堆
结构特点：任意一点具有大于/小于其左右子树中节点
基本操作：向上调整&向下调整
堆行为：
1.初始化堆：从非叶子节点开始逐个向下调整
2.出堆：堆顶节点和末尾节点交换，堆顶向下调整
3.移除堆中元素：该元素和末尾节点交换，先向下调整，如果无需向下则向上调整
4.修改堆中元素：和删除思路一致，省去交换，直接从修改位置向上向下调整
*/

package search_sort

// 调整：从上至下调整
// 注意参数传递，需要限制下界
func HeapSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	buildHeap(arr)

	// todo 根据需要使用堆
}

// 建堆 自下而上
func buildHeap(arr []int) {
	for iter := (len(arr) - 1) / 2; iter >= 0; iter-- {
		maxHeap(arr, iter, len(arr)-1)
	}
}

// 堆调整 自上而下
func maxHeap(arr []int, curRoot, end int) {
	left := 2*curRoot + 1
	right := 2*curRoot + 2
	max := curRoot

	if left < end && arr[left] > arr[max] {
		max = left
	}
	if right < end && arr[right] > arr[max] {
		max = right
	}

	if max != curRoot {
		arr[curRoot], arr[max] = arr[max], arr[curRoot]
		maxHeap(arr, max, end)
	}
}
