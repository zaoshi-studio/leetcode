package leetcode

import (
	"testing"
)

/**
Given two integers n and k, return all possible combinations of k numbers
chosen from the range [1, n].

 You may return the answer in any order.


 Example 1:


Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to
be the same combination.


 Example 2:


Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.



 Constraints:


 1 <= n <= 20
 1 <= k <= n


 Related Topics 回溯 👍 1324 👎 0

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
func combine(n int, k int) [][]int {

	var ans [][]int

	var backtrack func(idx int, buffer []int)
	backtrack = func(idx int, buffer []int) {
		if len(buffer) == k {
			tmp := make([]int, len(buffer))
			copy(tmp, buffer)
			ans = append(ans, tmp)
			return
		}

		for i := idx; i <= n; i++ {
			buffer = append(buffer, i)
			backtrack(i+1, buffer)
			buffer = buffer[:len(buffer)-1]
		}
	}

	backtrack(1, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinations(t *testing.T) {

}
