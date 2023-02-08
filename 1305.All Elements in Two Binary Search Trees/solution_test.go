package leetcode

import (
	"sort"
	"testing"
)

/**
Given two binary search trees root1 and root2, return a list containing all the
integers from both trees sorted in ascending order.


 Example 1:


Input: root1 = [2,1,4], root2 = [1,0,3]
Output: [0,1,1,2,3,4]


 Example 2:


Input: root1 = [1,null,8], root2 = [8,1]
Output: [1,1,8,8]



 Constraints:


 The number of nodes in each tree is in the range [0, 5000].
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 排序 👍 163 👎 0

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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	var ans []int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}

		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}

	inorder(root1)
	inorder(root2)

	sort.Ints(ans)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAllElementsInTwoBinarySearchTrees(t *testing.T) {

}
