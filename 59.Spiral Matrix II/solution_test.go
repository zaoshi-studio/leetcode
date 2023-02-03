package leetcode

import (
	"testing"
)

/**
Given a positive integer n, generate an n x n matrix filled with elements from 1
 to n² in spiral order.


 Example 1:


Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]


 Example 2:


Input: n = 1
Output: [[1]]



 Constraints:


 1 <= n <= 20


 Related Topics 数组 矩阵 模拟 👍 926 👎 0

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
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for k := range ans {
		ans[k] = make([]int, n)
	}

	var (
		top, left     int
		bottom, right = n - 1, n - 1
		idx           = 1
	)

	for idx <= n*n {
		// 从左到右
		for i := left; i <= right; i++ {
			ans[top][i] = idx
			idx++
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			ans[i][right] = idx
			idx++
		}
		right--

		// 从右到左
		for i := right; i >= left; i-- {
			ans[bottom][i] = idx
			idx++
		}
		bottom--

		// 从下到上
		for i := bottom; i >= top; i-- {
			ans[i][left] = idx
			idx++
		}
		left++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSpiralMatrixIi(t *testing.T) {

}
