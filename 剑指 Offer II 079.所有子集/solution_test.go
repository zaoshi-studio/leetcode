package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 位运算 数组 回溯 👍 58 👎 0

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
func subsets(nums []int) [][]int {
	var ans [][]int

	var n = len(nums)

	var backtrack func(idx int, buffer []int)
	backtrack = func(idx int, buffer []int) {
		if idx > n {
			return
		}

		tmp := make([]int, len(buffer))
		copy(tmp, buffer)
		ans = append(ans, tmp)

		for i := idx; i < n; i++ {
			buffer = append(buffer, nums[i])
			backtrack(i+1, buffer)
			buffer = buffer[:len(buffer)-1]
		}
	}

	backtrack(0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTVdhkn(t *testing.T) {

}
