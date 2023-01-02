package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数组 双指针 二分查找 👍 220 👎 0

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
func twoSum(nums []int, target int) []int {
	cur := 0

	for i, j := 0, len(nums)-1; i < j; {
		cur = nums[i] + nums[j]
		if cur > target {
			j--
		} else if cur < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHeWeiSdeLiangGeShuZiLcof(t *testing.T) {

}
