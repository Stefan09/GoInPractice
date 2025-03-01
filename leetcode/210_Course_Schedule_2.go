package leetcode

/* 课程表 拓扑排序
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：[0,1]
解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

示例 2：
输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
输出：[0,2,1,3]
解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

示例 3：
输入：numCourses = 1, prerequisites = []
输出：[0]
 */
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 拓扑排序
	// 1. 创建邻接表
	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}
	// 2. 创建入度表
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		inDegree[p[0]]++
	}
	// 3. 创建队列，将入度为0的节点入队
	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	// 4. 创建计数器，记录已经遍历过的节点数
	count := 0
	res := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		count++
		for _, v := range adj[node] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if count != numCourses {
		return []int{}
	}
	return res
}
