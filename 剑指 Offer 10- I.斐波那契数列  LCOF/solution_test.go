package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 记忆化搜索 数学 动态规划 👍 425 👎 0

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

func fib1(n int) int {
	if n <= 1 {
		return n
	}

	p, q, r := 0, 0, 1
	const mod int = 1e9 + 7

	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFeiBoNaQiShuLieLcof(t *testing.T) {

}
