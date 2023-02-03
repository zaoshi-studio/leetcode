package leetcode

import (
	"math"
	"testing"
)

/**
Given an array of positive integers nums and a positive integer target, return
the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.


 Example 1:


Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem
constraint.


 Example 2:


Input: target = 4, nums = [1,4,4]
Output: 1


 Example 3:


Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0



 Constraints:


 1 <= target <= 10⁹
 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁴



Follow up: If you have figured out the
O(n) solution, try coding another solution of which the time complexity is
O(n log(n)).

 Related Topics 数组 二分查找 前缀和 滑动窗口 👍 1545 👎 0

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
func minSubArrayLen(target int, nums []int) int {

	var (
		ans = math.MaxInt32
		sum int
	)

	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		// 收缩窗口
		// 为什么不要 i < j
		// 当 nums[x] == target 时，可能由于 i == j 导致 sum 得不到更新
		for sum >= target {
			ans = min(ans, j-i+1)
			sum -= nums[i]
			i++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumSizeSubarraySum(t *testing.T) {
	nums := []int{1, 4, 4}
	minSubArrayLen(4, nums)
}
