package leetcode

import (
	"sort"
	"testing"
)

/**
Write a method to compute all permutations of a string whose characters are not
necessarily unique. The list of permutations should not have duplicates.

 Example1:


 Input: S = "qqe"
 Output: ["eqq","qeq","qqe"]


 Example2:


 Input: S = "ab"
 Output: ["ab", "ba"]


 Note:


 All characters are English letters.
 1 <= S.length <= 9


 Related Topics 字符串 回溯 👍 82 👎 0

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
	raw := []byte(S)
	sort.Slice(raw, func(i, j int) bool {
		return raw[i] < raw[j]
	})

	var ans []string

	n := len(S)

	record := make([]bool, n)

	var backtrack func(idx int, buffer []byte)
	backtrack = func(idx int, buffer []byte) {
		if len(buffer) == n {
			ans = append(ans, string(buffer))
			return
		}

		for i := 0; i < n; i++ {
			// 如果当前字符已被使用，肯定要跳过
			// 如果当前字符与前一个字符相等，且都为访问过
			// 则第一个该字符已经完成了遍历，没参与了递归
			// 因为第一个字符参与的，后续相同的字符回都加入进去
			// 不能以相同的第2~n个字符重新开始去填充
			if record[i] || i > 0 && raw[i] == raw[i-1] && !record[i-1] && !record[i] {
				continue
			}

			buffer = append(buffer, raw[i])
			record[i] = true
			backtrack(i+1, buffer)
			record[i] = false
			buffer = buffer[:len(buffer)-1]
		}
	}

	backtrack(0, []byte{})
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPermutationIiLcci(t *testing.T) {

}
