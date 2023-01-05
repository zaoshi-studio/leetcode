package leetcode

import (
	"fmt"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 栈 树 二叉搜索树 递归 二叉树 单调栈 👍 643 👎 0

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
func verifyPostorder(postorder []int) bool {

	var traverse func(postStart, postEnd int) bool
	traverse = func(postStart, postEnd int) bool {

		if postStart >= postEnd {
			return true
		}

		var idx int

		nodeVal := postorder[postEnd]

		// 检查右子树
		for idx = postEnd - 1; idx > postStart && postorder[idx] > nodeVal; idx-- {
			if postorder[idx] < nodeVal {
				return false
			}
		}

		// 检查左子树
		for i := postStart; postStart < idx && i <= idx; i++ {
			if postorder[i] > nodeVal {
				return false
			}
		}

		return traverse(postStart, idx) && traverse(idx+1, postEnd-1)
	}

	return traverse(0, len(postorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaSouSuoShuDeHouXuBianLiXuLieLcof(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(verifyPostorder(nums))
}
