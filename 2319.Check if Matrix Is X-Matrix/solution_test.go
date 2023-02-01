package leetcode

import (
	"testing"
)

/**
A square matrix is said to be an X-Matrix if both of the following conditions
hold:


 All the elements in the diagonals of the matrix are non-zero.
 All other elements are 0.


 Given a 2D integer array grid of size n x n representing a square matrix,
return true if grid is an X-Matrix. Otherwise, return false.


 Example 1:


Input: grid = [[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
Output: true
Explanation: Refer to the diagram above.
An X-Matrix should have the green elements (diagonals) be non-zero and the red
elements be 0.
Thus, grid is an X-Matrix.


 Example 2:


Input: grid = [[5,7,0],[0,3,1],[0,5,0]]
Output: false
Explanation: Refer to the diagram above.
An X-Matrix should have the green elements (diagonals) be non-zero and the red
elements be 0.
Thus, grid is not an X-Matrix.



 Constraints:


 n == grid.length == grid[i].length
 3 <= n <= 100
 0 <= grid[i][j] <= 10⁵


 Related Topics 数组 矩阵 👍 27 👎 0

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
func checkXMatrix(grid [][]int) bool {
	if nil == grid {
		return false
	}

	var (
		n = len(grid)
	)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCheckIfMatrixIsXMatrix(t *testing.T) {

}
