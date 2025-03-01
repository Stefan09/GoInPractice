package main

import (
	"fmt"
	"math"
)

var tuopu []int // 用于存储拓扑排序结果

// dfsVisit 是递归访问图的辅助函数
func dfsVisit(graph [][]int, node int, visit []int, father []int) {
	n := len(graph) // 获取图的节点数
	visit[node] = 1 // 标记当前节点为已访问（状态1）

	// 遍历所有节点
	for i := 0; i < n; i++ {
		// 如果节点i与当前节点node有边且不是同一个节点
		if i != node && graph[node][i] != math.MaxInt32 {
			// 如果节点i已经被访问且不是当前节点的父节点，则发现一个环
			if visit[i] == 1 && i != father[node] {
				tmp := node // 从当前节点开始回溯
				fmt.Print("cycle: ")
				for tmp != i { // 输出环的路径
					fmt.Print(tmp, "->")
					tmp = father[tmp]
				}
				fmt.Println(tmp) // 输出环的起始节点
			} else if visit[i] == 0 { // 如果节点i未被访问，则递归访问节点i
				father[i] = node                  // 设置节点i的父节点为当前节点
				dfsVisit(graph, i, visit, father) // 递归调用
			}
		}
	}

	visit[node] = 2             // 标记当前节点为完全访问（状态2）
	tuopu = append(tuopu, node) // 将当前节点加入拓扑排序结果
}

func main() {
	input := [][]int{{1, 0}, {2, 0}, {3, 1}, {0, 3}}
	graph := make([][]int, 4)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 4)
	}
	for i := 0; i < len(input); i++ {
		graph[input[i][1]][input[i][0]] = 1
	}
	fmt.Println(graph)

	visited, path := make([]int, 4), make([]int, 0)
	dfs(graph, 0, visited, &path)
}

func dfs(graph [][]int, node int, visited []int, path *[]int) {
	visited[node] = 1
	*path = append(*path, node)
	defer func() {
		*path = (*path)[:len(*path)-1]
		visited[node] = 0
	}()

	for i := 0; i < 4; i++ {
		if graph[node][i] == 1 {
			if visited[i] == 1 {
				fmt.Println("has circle", path, i)
			} else {
				dfs(graph, i, visited, path)
			}
		}
	}
}
