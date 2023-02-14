package leetcode

import (
	"testing"
)

/**
Given the root of a perfect binary tree, reverse the node values at each odd
level of the tree.


 For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then
it should become [18,29,11,7,4,3,1,2].


 Return the root of the reversed tree.

 A binary tree is perfect if all parent nodes have two children and all leaves
are on the same level.

 The level of a node is the number of edges along the path between it and the
root node.


 Example 1:


Input: root = [2,3,5,8,13,21,34]
Output: [2,5,3,8,13,21,34]
Explanation:
The tree has only one odd level.
The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.


 Example 2:


Input: root = [7,13,11]
Output: [7,11,13]
Explanation:
The nodes at level 1 are 13, 11, which are reversed and become 11, 13.


 Example 3:


Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
Explanation:
The odd levels have non-zero values.
The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1
 after the reversal.



 Constraints:


 The number of nodes in the tree is in the range [1, 2¹⁴].
 0 <= Node.val <= 10⁵
 root is a perfect binary tree.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 20 👎 0

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
func reverseOddLevels(root *TreeNode) *TreeNode {

	if nil == root {
		return root
	}

	var reverse = true

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]

		if reverse {
			for i, j := 0, len(queue)-1; i < j; {
				queue[i].Val, queue[j].Val = queue[j].Val, queue[i].Val
				i++
				j--
			}
		}

		reverse = !reverse
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseOddLevelsOfBinaryTree(t *testing.T) {

}
