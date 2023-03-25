package leetcode

import (
	"sort"
	"testing"
)

/**
Given a collection of numbers, nums, that might contain duplicates, return all
possible unique permutations in any order.


 Example 1:


Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]


 Example 2:


Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]



 Constraints:


 1 <= nums.length <= 8
 -10 <= nums[i] <= 10


 Related Topics 数组 回溯 👍 1323 👎 0

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
func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)

	var (
		ans  [][]int
		n    = len(nums)
		used = make(map[int]bool, n)
	)

	var backtrack func(idx int, buffer []int)
	backtrack = func(idx int, buffer []int) {
		if len(buffer) == n {
			tmp := make([]int, n)
			copy(tmp, buffer)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] || used[i] {
				continue
			}
			used[i] = true
			buffer = append(buffer, nums[i])
			backtrack(i+1, buffer)
			buffer = buffer[:len(buffer)-1]
			used[i] = false
		}
	}

	backtrack(0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationsIi(t *testing.T) {

}
