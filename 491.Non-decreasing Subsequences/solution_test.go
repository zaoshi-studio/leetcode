package leetcode

import (
	"testing"
)

/**
Given an integer array nums, return all the different possible non-decreasing
subsequences of the given array with at least two elements. You may return the
answer in any order.


 Example 1:


Input: nums = [4,6,7,7]
Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]


 Example 2:


Input: nums = [4,4,3,2,1]
Output: [[4,4]]



 Constraints:


 1 <= nums.length <= 15
 -100 <= nums[i] <= 100


 Related Topics 位运算 数组 哈希表 回溯 👍 623 👎 0

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
func findSubsequences(nums []int) [][]int {
	var (
		ans [][]int
		n   = len(nums)
		//record = make(map[int]bool)
	)

	var backtrack func(idx int, buffer []int)
	backtrack = func(idx int, buffer []int) {
		if len(buffer) >= 2 {
			tmp := make([]int, len(buffer))
			copy(tmp, buffer)
			ans = append(ans, tmp)
		}

		// layer record
		// guarantee the layer is a non-decreasing sequence
		used := make(map[int]bool, len(nums))

		for i := idx; i < n; i++ {

			if used[nums[i]] {
				continue
			}

			// if nums[i] is no less than the last one in buffer, then add it
			// if buffer is empty, add it too
			if len(buffer) == 0 || buffer[len(buffer)-1] <= nums[i] {
				used[nums[i]] = true
				buffer = append(buffer, nums[i])

				backtrack(i+1, buffer)

				buffer = buffer[:len(buffer)-1]
			}
		}
	}

	backtrack(0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNonDecreasingSubsequences(t *testing.T) {
	input := []int{1, 2, 3, 1, 1}
	findSubsequences(input)
}
