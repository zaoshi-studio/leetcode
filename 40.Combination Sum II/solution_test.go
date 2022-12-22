package leetcode

import (
	"sort"
	"testing"
)

/**
Given a collection of candidate numbers (candidates) and a target number (
target), find all unique combinations in candidates where the candidate numbers sum
to target.

 Each number in candidates may only be used once in the combination.

 Note: The solution set must not contain duplicate combinations.


 Example 1:


Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]


 Example 2:


Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]



 Constraints:


 1 <= candidates.length <= 100
 1 <= candidates[i] <= 50
 1 <= target <= 30


 Related Topics 数组 回溯 👍 1189 👎 0

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
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	var ans [][]int
	//used := make(map[int]bool, len(candidates))

	var backtrack func(idx, sum int, buffer []int)
	backtrack = func(idx, sum int, buffer []int) {
		if idx > len(candidates) || sum > target {
			return
		}

		if sum == target {
			temp := make([]int, len(buffer))
			copy(temp, buffer)
			ans = append(ans, temp)
			return
		}

		for i := idx; i < len(candidates); i++ {

			// 注意此处是 i > idx
			// candidates[i] == candidates[i-1] 保证了该 for 循环中不出现重复
			// i > idx 保证当前的 candidates[idx] 不受前一个影响
			// 即使 candidates[idx] == candidates[idx-1] 也能选择当前元素
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			//used[i] = true
			sum += candidates[i]
			buffer = append(buffer, candidates[i])
			backtrack(i+1, sum, buffer)
			sum -= candidates[i]
			buffer = buffer[:len(buffer)-1]
			//used[i] = false
		}

	}

	backtrack(0, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCombinationSumIi(t *testing.T) {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	combinationSum2(nums, target)
}
