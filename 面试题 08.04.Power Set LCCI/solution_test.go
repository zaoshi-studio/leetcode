package leetcode

import (
	"testing"
)

/**
Write a method to return all subsets of a set. The elements in a set are
pairwise distinct.

 Note: The result set should not contain duplicated subsets.

 Example:


 Input:  nums = [1,2,3]
 Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]


 Related Topics 位运算 数组 回溯 👍 102 👎 0

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

func TestPowerSetLcci(t *testing.T) {

}
