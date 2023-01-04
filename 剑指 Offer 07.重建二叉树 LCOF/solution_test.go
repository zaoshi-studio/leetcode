package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 955 👎 0

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

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	record := make(map[int]int, len(preorder))
	for k, v := range inorder {
		record[v] = k
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {

		if preStart > preEnd {
			return nil
		}

		idx := record[preorder[preStart]]

		leftSize := idx - inStart

		left := build(preStart+1, preStart+leftSize, inStart, idx-1)
		right := build(preStart+leftSize+1, preEnd, idx+1, inEnd)

		node := &TreeNode{
			Val:   preorder[preStart],
			Left:  left,
			Right: right,
		}

		return node
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZhongJianErChaShuLcof(t *testing.T) {

}
