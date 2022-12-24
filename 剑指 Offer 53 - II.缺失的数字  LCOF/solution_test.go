package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 位运算 数组 哈希表 数学 二分查找 👍 328 👎 0

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
func missingNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestQueShiDeShuZiLcof(t *testing.T) {

}
