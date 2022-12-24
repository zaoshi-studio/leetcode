package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 数组 哈希表 分治 计数 排序 👍 335 👎 0

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
func majorityElement(nums []int) int {
	sort.Ints(nums)

	candidate := nums[0]
	vote := 0

	for _, v := range nums {

		if vote == 0 {
			candidate = v
		}

		if candidate == v {
			vote++
		} else {
			vote--
		}

	}

	return candidate
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShuZuZhongChuXianCiShuChaoGuoYiBanDeShuZiLcof(t *testing.T) {

}
