package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 栈 数组 模拟 👍 396 👎 0

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
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	var stack []int
	p := 0
	for _, v := range pushed {
		stack = append(stack, v)

		for len(stack) > 0 && stack[len(stack)-1] == popped[p] {
			stack = stack[:len(stack)-1]
			p++
		}
	}

	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZhanDeYaRuDanChuXuLieLcof(t *testing.T) {

}
