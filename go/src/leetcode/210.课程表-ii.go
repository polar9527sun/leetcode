/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 *
 * https://leetcode-cn.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (46.86%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    23.1K
 * Total Submissions: 48.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 n 门课需要选，记为 0 到 n-1。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
 *
 * 给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
 *
 * 可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: [0,1]
 * 解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 *
 * 示例 2:
 *
 * 输入: 4, [[1,0],[2,0],[3,1],[3,2]]
 * 输出: [0,1,2,3] or [0,2,1,3]
 * 解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 *
 *
 * 说明:
 *
 *
 * 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 *
 *
 * 提示:
 *
 *
 * 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
 * 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
 *
 * 拓扑排序也可以通过 BFS 完成。
 *
 *
 *
 */
package main

// @lc code=start
// BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indegs := make([]int, numCourses)
	var result []int
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indegs[info[0]]++
	}
	dqueue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegs[i] == 0 {
			dqueue = append(dqueue, i)
		}
	}

	for len(dqueue) != 0 {
		cur := dqueue[0]
		dqueue = dqueue[1:]
		for _, neighbor := range edges[cur] {
			indegs[neighbor]--
			if indegs[neighbor] == 0 {
				dqueue = append(dqueue, neighbor)
			}
		}
		result = append(result, cur)
	}

	if len(result) != numCourses {
		return []int{}
	}
	return result
}

// DFS
// type state int
//
// const (
// 	undiscovered state = iota
// 	discovered
// 	visited
// )

// func findOrder(numCourses int, prerequisites [][]int) []int {
// 	edges := make([][]int, numCourses)
// 	visitedNodes := make([]state, numCourses)
// 	var result []int
// 	var dfs func(cur int) bool
// 	dfs = func(cur int) bool {
// 		visitedNodes[cur] = discovered
// 		for _, neighbor := range edges[cur] {
// 			if visitedNodes[neighbor] == undiscovered {
// 				if !dfs(neighbor) {
// 					return false
// 				}
// 			} else if visitedNodes[neighbor] == discovered {
// 				// 	not DAG
// 				return false
// 			}
// 		}
// 		visitedNodes[cur] = visited
// 		result = append(result, cur)
// 		return true
// 	}
//
// 	// establish all edges
// 	for _, info := range prerequisites {
// 		// [1,0] :   info[1]-->info[0],  0-->1
// 		edges[info[1]] = append(edges[info[1]], info[0])
// 	}
//
// 	for i := 0; i < numCourses; i++ {
// 		if visitedNodes[i] == undiscovered {
// 			// not DAG
// 			if !dfs(i) {
// 				return []int{}
// 			}
// 		}
// 	}
// 	for i := 0; i < len(result)/2; i++ {
// 		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
// 	}
// 	return result
// }

// @lc code=end
