package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 字符串 动态规划 回溯 👍 61 👎 0

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
func generateParenthesis(n int) []string {

	var (
		ans []string
	)

	var backtrack func(left, right int, buffer []byte)
	backtrack = func(left, right int, buffer []byte) {

		if left > right || left < 0 || right < 0 {
			return
		}

		if left == 0 && right == 0 {
			tmp := make([]byte, len(buffer))
			copy(tmp, buffer)
			ans = append(ans, string(tmp))
			return
		}

		// use code position to make priority
		// prefer left bracket
		buffer = append(buffer, '(')
		backtrack(left-1, right, buffer)

		buffer[len(buffer)-1] = ')'
		backtrack(left, right-1, buffer)

		// reset buffer
		buffer = buffer[:len(buffer)-1]
	}

	backtrack(n, n, []byte{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIDBivT(t *testing.T) {

}
