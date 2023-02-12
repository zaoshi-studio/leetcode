package leetcode

import (
	"testing"
)

/**
Given a binary tree


struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}


 Populate each next pointer to point to its next right node. If there is no
next right node, the next pointer should be set to NULL.

 Initially, all next pointers are set to NULL.


 Example 1:


Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should
populate each next pointer to point to its next right node, just like in Figure B.
The serialized output is in level order as connected by the next pointers, with '#
' signifying the end of each level.


 Example 2:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 6000].
 -100 <= Node.val <= 100



 Follow-up:


 You may only use constant extra space.
 The recursive approach is fine. You may assume implicit stack space does not
count as extra space for this problem.


 Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 680 👎 0

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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {

	if nil == root {
		return nil
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		size := len(queue)

		var pre *Node

		for i := 0; i < size; i++ {
			node := queue[i]

			// 每层从右向左遍历即可
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			node.Next = pre
			pre = node
		}
		queue = queue[size:]

	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPopulatingNextRightPointersInEachNodeIi(t *testing.T) {

}
