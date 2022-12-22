package leetcode

import (
	"testing"
)

/**
Given a sorted array of distinct integers and a target value, return the index
if the target is found. If not, return the index where it would be if it were
inserted in order.

 You must write an algorithm with O(log n) runtime complexity.


 Example 1:


Input: nums = [1,3,5,6], target = 5
Output: 2


 Example 2:


Input: nums = [1,3,5,6], target = 2
Output: 1


 Example 3:


Input: nums = [1,3,5,6], target = 7
Output: 4



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums contains distinct values sorted in ascending order.
 -10⁴ <= target <= 10⁴


 Related Topics 数组 二分查找 👍 1823 👎 0

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
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {

		mid := (left + right) / 2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInsertPosition(t *testing.T) {

}
