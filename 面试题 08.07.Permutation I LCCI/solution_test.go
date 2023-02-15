package leetcode

import (
	"testing"
)

/**
Write a method to compute all permutations of a string of unique characters.

 Example1:


 Input: S = "qwe"
 Output: ["qwe", "qew", "wqe", "weq", "ewq", "eqw"]


 Example2:


 Input: S = "ab"
 Output: ["ab", "ba"]


 Note:


 All charaters are English letters.
 1 <= S.length <= 9


 Related Topics 字符串 回溯 👍 84 👎 0

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
func permutation(S string) []string {
	var ans []string

	record := make(map[int]bool, len(S))

	var n = len(S)

	var backtrack func(buffer []rune)
	backtrack = func(buffer []rune) {
		if len(buffer) == n {
			tmp := make([]rune, n)
			copy(tmp, buffer)
			ans = append(ans, string(tmp))
			return
		}

		for k, v := range S {
			// if the elem at nums[k] has been added to buffer, skip it
			if record[k] {
				continue
			}

			record[k] = true
			backtrack(append(buffer, v))
			record[k] = false
		}
	}

	backtrack([]rune{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationILcci(t *testing.T) {

}
