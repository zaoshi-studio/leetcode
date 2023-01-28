package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 记忆化搜索 数学 动态规划 👍 355 👎 0

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
func numWays(n int) int {

	const mod = 1e9 + 7

	f1, f2, f3 := 0, 0, 1

	for i := 1; i <= n; i++ {
		f1 = f2
		f2 = f3
		f3 = (f1 + f2) % mod
	}

	return f3
}

//leetcode submit region end(Prohibit modification and deletion)

func TestQingWaTiaoTaiJieWenTiLcof(t *testing.T) {

}
