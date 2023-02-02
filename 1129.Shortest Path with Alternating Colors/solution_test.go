package leetcode

import (
	"testing"
)

/**
You are given an integer n, the number of nodes in a directed graph where the
nodes are labeled from 0 to n - 1. Each edge is red or blue in this graph, and
there could be self-edges and parallel edges.

 You are given two arrays redEdges and blueEdges where:


 redEdges[i] = [ai, bi] indicates that there is a directed red edge from node
ai to node bi in the graph, and
 blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node
uj to node vj in the graph.


 Return an array answer of length n, where each answer[x] is the length of the
shortest path from node 0 to node x such that the edge colors alternate along
the path, or -1 if such a path does not exist.


 Example 1:


Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
Output: [0,1,-1]


 Example 2:


Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
Output: [0,1,-1]



 Constraints:


 1 <= n <= 100
 0 <= redEdges.length, blueEdges.length <= 400
 redEdges[i].length == blueEdges[j].length == 2
 0 <= ai, bi, uj, vj < n


 Related Topics 广度优先搜索 图 👍 166 👎 0

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)
//func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
//
//	ans := make([]int, n)
//
//	//var redUsed [][]int
//
//	var bfs func(start, idx int, isRed bool)
//	bfs = func(start, idx int, isRed bool) {
//		// 在当前颜色路径中查找是否有以当前结点开始的有向边
//		var (
//			curEdges, nextEdges [][]int
//		)
//
//		nextStart := make(map[int]struct{})
//
//		if isRed {
//			curEdges = redEdges
//			nextEdges = blueEdges
//		} else {
//			curEdges = blueEdges
//			nextEdges = redEdges
//		}
//
//		// 遍历当前路径
//		for _, v := range curEdges {
//			// 找到以当前结点开始的有向边
//			if v[0] == start {
//				nextStart[v[1]] = struct{}{}
//			}
//		}
//
//		// 遍历下一个颜色要访问的路径
//		for _, v := range nextEdges {
//			// 有从当前节点到下一个颜色不同结点的路径
//			if _, ok := nextStart[v[0]]; ok {
//				ans[idx] = 1
//				idx++
//				bfs(v[0], idx, !isRed)
//			}
//		}
//
//	}
//
//	bfs(0, 1, true)
//
//	return ans
//}

func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	g := make([][]pair, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	vis := make([][2]bool, n)
	vis[0] = [2]bool{true, true}
	q := []pair{{0, 0}, {0, 1}}
	for level := 0; len(q) > 0; level++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			x := p.x
			if dis[x] < 0 {
				dis[x] = level
			}
			for _, to := range g[x] {
				if to.color != p.color && !vis[to.x][to.color] {
					vis[to.x][to.color] = true
					q = append(q, to)
				}
			}
		}
	}
	return dis
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShortestPathWithAlternatingColors(t *testing.T) {
	n := 3
	redEdges := [][]int{{0, 1}, {0, 2}}
	blueEdges := [][]int{{1, 0}}
	shortestAlternatingPaths(n, redEdges, blueEdges)
}
