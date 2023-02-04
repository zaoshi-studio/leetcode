package leetcode

import (
	"sort"
	"testing"
)

/**
You are given an integer array coins of length n which represents the n coins
that you own. The value of the iᵗʰ coin is coins[i]. You can make some value x if
you can choose some of your n coins such that their values sum up to x.

 Return the maximum number of consecutive integer values that you can make with
your coins starting from and including 0.

 Note that you may have multiple coins of the same value.


 Example 1:


Input: coins = [1,3]
Output: 2
Explanation: You can make the following values:
- 0: take []
- 1: take [1]
You can make 2 consecutive integer values starting from 0.

 Example 2:


Input: coins = [1,1,1,4]
Output: 8
Explanation: You can make the following values:
- 0: take []
- 1: take [1]
- 2: take [1,1]
- 3: take [1,1,1]
- 4: take [4]
- 5: take [4,1]
- 6: take [4,1,1]
- 7: take [4,1,1,1]
You can make 8 consecutive integer values starting from 0.

 Example 3:


Input: nums = [1,4,10,3,1]
Output: 20


 Constraints:


 coins.length == n
 1 <= n <= 4 * 10⁴
 1 <= coins[i] <= 4 * 10⁴


 Related Topics 贪心 数组 👍 115 👎 0

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
func getMaximumConsecutive(coins []int) int {

	sort.Ints(coins)

	var ans = 1

	for _, v := range coins {
		if v > ans {
			break
		}

		// ans 每次加上当前未使用的最小数字，形成一个 [0, ans] 的区间
		// 在 [0, ans] 的区间中所有数字都能被构造
		ans += v
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumNumberOfConsecutiveValuesYouCanMake(t *testing.T) {

}
