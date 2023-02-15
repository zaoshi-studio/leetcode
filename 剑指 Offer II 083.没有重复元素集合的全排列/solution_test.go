package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数组 回溯 👍 45 👎 0

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
func permute(nums []int) [][]int {

	var ans [][]int

	record := make(map[int]bool, len(nums))

	var n = len(nums)

	var backtrack func(buffer []int)
	backtrack = func(buffer []int) {
		if len(buffer) == n {
			tmp := make([]int, n)
			copy(tmp, buffer)
			ans = append(ans, tmp)
			return
		}

		for k, v := range nums {
			// if the elem at nums[k] has been added to buffer, skip it
			if record[k] {
				continue
			}

			record[k] = true
			backtrack(append(buffer, v))
			record[k] = false
		}
	}

	backtrack([]int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestVvJkup(t *testing.T) {
	nums := []int{1, 2, 3}
	permute(nums)
}
