package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数组 哈希表 排序 👍 1024 👎 0

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
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for k, v := range nums {
		if k > 0 && nums[k-1] == v {
			return v
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShuZuZhongZhongFuDeShuZiLcof(t *testing.T) {

}
