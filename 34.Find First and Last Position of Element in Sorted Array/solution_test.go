package leetcode

import (
	"testing"
)

/**
Given an array of integers nums sorted in non-decreasing order, find the
starting and ending position of a given target value.

 If target is not found in the array, return [-1, -1].

 You must write an algorithm with O(log n) runtime complexity.


 Example 1:
 Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

 Example 2:
 Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

 Example 3:
 Input: nums = [], target = 0
Output: [-1,-1]


 Constraints:


 0 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 nums is a non-decreasing array.
 -10⁹ <= target <= 10⁹


 Related Topics 数组 二分查找 👍 2065 👎 0

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
func searchRange(nums []int, target int) []int {

	v1 := binarySearch(nums, target, true)
	v2 := binarySearch(nums, target, false)

	if v1 == len(nums) || nums[v1] != target {
		return []int{-1, -1}
	}
	return []int{v1, v2}
}

func binarySearch(nums []int, target int, flag bool) int {

	left, right := 0, len(nums)-1
	mid := 0

	// 左闭右闭区间
	for left < right {
		mid = (left + right) / 2

		if nums[mid] < target && flag {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {

}
