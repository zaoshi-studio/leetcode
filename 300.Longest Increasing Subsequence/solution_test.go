package leetcode

import (
	"testing"
)

/**
Given an integer array nums, return the length of the longest strictly
increasing subsequence.


 Example 1:


Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the
length is 4.


 Example 2:


Input: nums = [0,1,0,3,2,3]
Output: 4


 Example 3:


Input: nums = [7,7,7,7,7,7,7]
Output: 1



 Constraints:


 1 <= nums.length <= 2500
 -10⁴ <= nums[i] <= 10⁴



 Follow up: Can you come up with an algorithm that runs in O(n log(n)) time
complexity?

 Related Topics 数组 二分查找 动态规划 👍 2939 👎 0

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
func lengthOfLIS(nums []int) int {

	ans := 1

	dp := make([]int, len(nums))

	for k := range nums {
		dp[k] = 1
	}

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[i]+1, dp[j])
				ans = max(ans, dp[j])
			}
		}
	}

	return ans
}

func lengthOfLIS2(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var ans int

	dp := make([]int, len(nums))

	for k := range nums {
		dp[k] = 1
	}

	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestIncreasingSubsequence(t *testing.T) {

}
