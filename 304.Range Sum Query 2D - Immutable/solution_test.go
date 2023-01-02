package leetcode

import (
	"testing"
)

/**
Given a 2D matrix matrix, handle multiple queries of the following type:


 Calculate the sum of the elements of matrix inside the rectangle defined by
its upper left corner (row1, col1) and lower right corner (row2, col2).


 Implement the NumMatrix class:


 NumMatrix(int[][] matrix) Initializes the object with the integer matrix
matrix.
 int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the
elements of matrix inside the rectangle defined by its upper left corner (row1,
col1) and lower right corner (row2, col2).


 You must design an algorithm where sumRegion works on O(1) time complexity.


 Example 1:


Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3,
 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Output
[null, 8, 11, 12]


Explanation
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0,
 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)



 Constraints:


 m == matrix.length
 n == matrix[i].length
 1 <= m, n <= 200
 -10⁴ <= matrix[i][j] <= 10⁴
 0 <= row1 <= row2 < m
 0 <= col1 <= col2 < n
 At most 10⁴ calls will be made to sumRegion.


 Related Topics 设计 数组 矩阵 前缀和 👍 451 👎 0

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
type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])

	preSum := make([][]int, row+1)
	for i := range preSum {
		preSum[i] = make([]int, col+1)
	}

	// 每个前缀和计算 [1, row] 和 [1, col] 包裹起来的 matrix 中值的总和
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}

	return NumMatrix{
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 加的是重合部分，减去时被多减了一次
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRangeSumQuery2dImmutable(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	s := Constructor(matrix)

	s.SumRegion(2, 1, 4, 3)
}
