package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数组 回溯 👍 43 👎 0

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
func combinationSum(candidates []int, target int) [][]int {

	// sort the array, in order to make the same number together
	sort.Ints(candidates)

	var ans [][]int

	n := len(candidates)

	var backtrack func(idx, preSum int, buffer []int)
	backtrack = func(idx, preSum int, buffer []int) {
		if idx > n || preSum > target {
			return
		}

		if preSum == target {
			tmp := make([]int, len(buffer))
			copy(tmp, buffer)
			ans = append(ans, tmp)
			return
		}

		for i := idx; i < n; i++ {
			// every time start loop from the idx
			// if there are same numbers, need to skip it
			// this mean we can select the number at the same position sometimes
			// but not the same number at different position
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			buffer = append(buffer, candidates[i])
			backtrack(i, preSum+candidates[i], buffer)
			buffer = buffer[:len(buffer)-1]
		}

	}

	backtrack(0, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestYgoe9J(t *testing.T) {

}
