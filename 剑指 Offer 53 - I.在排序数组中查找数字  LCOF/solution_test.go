package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 数组 二分查找 👍 383 👎 0

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
func search(nums []int, target int) int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return right - left + 1
	}
	return 0
}

func binarySearch(nums []int, target int, flag bool) int {
	left, right := 0, len(nums)-1
	ans := len(nums)

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > target || (flag && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZaiPaiXuShuZuZhongChaZhaoShuZiLcof(t *testing.T) {

}
