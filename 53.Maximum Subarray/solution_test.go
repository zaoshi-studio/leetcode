package leetcode

import (
	"testing"
)

/**
Given an integer array nums, find the subarray with the largest sum, and return
its sum.


 Example 1:


Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.


 Example 2:


Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.


 Example 3:


Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.



 Constraints:


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴



 Follow up: If you have figured out the O(n) solution, try coding another
solution using the divide and conquer approach, which is more subtle.

 Related Topics 数组 分治 动态规划 👍 5665 👎 0

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
func maxSubArray(nums []int) int {

	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		ans = max(ans, nums[i])
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

func TestMaximumSubarray(t *testing.T) {

}
