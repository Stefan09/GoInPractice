package leetcode

/* 课程表 拓扑排序
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 拓扑排序
	// 1. 创建邻接表
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
	}
	// 2. 创建入度数组
	in := make([]int, numCourses)
	for _, p := range prerequisites {
		in[p[0]]++
	}
	// 3. 创建队列，将入度为0的节点入队
	queue := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	// 4. 创建计数器，记录已经遍历过的节点数
	count := 0
	// 5. 遍历队列，将队列中的节点出队，并将该节点的邻接节点的入度减1，如果入度减1后为0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++
		for _, v := range adj[node] {
			in[v]--
			if in[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return count == numCourses
}
