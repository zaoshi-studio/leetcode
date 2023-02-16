package leetcode

import (
	"testing"
)

/**
Find all valid combinations of k numbers that sum up to n such that the
following conditions are true:


 Only numbers 1 through 9 are used.
 Each number is used at most once.


 Return a list of all possible valid combinations. The list must not contain
the same combination twice, and the combinations may be returned in any order.


 Example 1:


Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

 Example 2:


Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.


 Example 3:


Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2
+3+4 = 10 and since 10 > 1, there are no valid combination.



 Constraints:


 2 <= k <= 9
 1 <= n <= 60


 Related Topics 数组 回溯 👍 619 👎 0

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
func combinationSum3(k int, n int) [][]int {
	var ans [][]int

	var backtrack func(idx, preSum int, buffer []int)
	backtrack = func(idx, preSum int, buffer []int) {
		if len(buffer) > k || preSum > n {
			return
		}

		if preSum == n && len(buffer) == k {
			tmp := make([]int, len(buffer))
			copy(tmp, buffer)
			ans = append(ans, tmp)
			return
		}

		for i := idx; i <= 9; i++ {
			buffer = append(buffer, i)
			backtrack(i+1, preSum+i, buffer)
			buffer = buffer[:len(buffer)-1]
		}
	}

	backtrack(1, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinationSumIii(t *testing.T) {

}
