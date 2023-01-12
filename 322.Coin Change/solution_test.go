package leetcode

import (
	"math"
	"testing"
)

/**
You are given an integer array coins representing coins of different
denominations and an integer amount representing a total amount of money.

 Return the fewest number of coins that you need to make up that amount. If
that amount of money cannot be made up by any combination of the coins, return -1.

 You may assume that you have an infinite number of each kind of coin.


 Example 1:


Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1


 Example 2:


Input: coins = [2], amount = 3
Output: -1


 Example 3:


Input: coins = [1], amount = 0
Output: 0



 Constraints:


 1 <= coins.length <= 12
 1 <= coins[i] <= 2³¹ - 1
 0 <= amount <= 10⁴


 Related Topics 广度优先搜索 数组 动态规划 👍 2257 👎 0

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
func coinChange1(coins []int, amount int) int {

	// record min op for evey amount
	record := make(map[int]int, amount+1)

	var dp func(remain int) int
	dp = func(remain int) int {

		if remain == 0 {
			return 0
		}

		if remain < 0 {
			return -1
		}

		if rop, ok := record[remain]; ok {
			return rop
		}

		var (
			minop = math.MaxInt
		)

		for _, v := range coins {
			subop := dp(remain - v)
			if subop == -1 {
				continue
			}
			minop = min(subop, minop)
		}

		if minop == math.MaxInt {
			record[remain] = -1
			return -1
		}

		record[remain] = minop + 1

		return minop + 1
	}

	dp(amount)

	return record[amount]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		// important!
		// can not set dp[i] default value as math.MathInt
		// when math.MathInt + 1 will overflow convert to math.MinInt
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i] = min(dp[i-coin]+1, dp[i])
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCoinChange(t *testing.T) {
	coins := []int{2}
	amount := 3

	coinChange(coins, amount)
}
