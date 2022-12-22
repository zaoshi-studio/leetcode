package leetcode

import (
	"testing"
)

/**
Given a string containing digits from 2-9 inclusive, return all possible letter
combinations that the number could represent. Return the answer in any order.

 A mapping of digits to letters (just like on the telephone buttons) is given
below. Note that 1 does not map to any letters.


 Example 1:


Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]


 Example 2:


Input: digits = ""
Output: []


 Example 3:


Input: digits = "2"
Output: ["a","b","c"]



 Constraints:


 0 <= digits.length <= 4
 digits[i] is a digit in the range ['2', '9'].


 Related Topics 哈希表 字符串 回溯 👍 2231 👎 0

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
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	keyboard := map[byte][]byte{}
	keyboard[49] = []byte{}
	keyboard[50] = []byte("abc")
	keyboard[51] = []byte("def")
	keyboard[52] = []byte("ghi")
	keyboard[53] = []byte("jkl")
	keyboard[54] = []byte("mno")
	keyboard[55] = []byte("pqrs")
	keyboard[56] = []byte("tuv")
	keyboard[57] = []byte("wxyz")

	var ans []string
	var buffer []byte

	var backtrack func(out, in int)

	backtrack = func(out, in int) {
		if len(buffer) == len(digits) {
			ans = append(ans, string(buffer))
			return
		}
		if out >= len(digits) {
			return
		}
		// 内部可选数字循环完成, 就跳转到外部的下一个数字
		if in >= len(keyboard[digits[out]]) {
			backtrack(out+1, 0)
			return
		}

		// 选择实际上是选择 out 中的第 in 个
		buffer = append(buffer, keyboard[digits[out]][in])
		// out 中选完之后，要跳转到下一个数字
		backtrack(out+1, 0)
		buffer = buffer[:len(buffer)-1]
		backtrack(out, in+1)

	}

	backtrack(0, 0)
	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {

}
