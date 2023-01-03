package leetcode

import (
	"testing"
)

/**
You are given a perfect binary tree where all leaves are on the same level, and
every parent has two children. The binary tree has the following definition:


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


Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function
should populate each next pointer to point to its next right node, just like in
Figure B. The serialized output is in level order as connected by the next pointers,
with '#' signifying the end of each level.


 Example 2:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 2¹² - 1].
 -1000 <= Node.val <= 1000



 Follow-up:


 You may only use constant extra space.
 The recursive approach is fine. You may assume implicit stack space does not
count as extra space for this problem.


 Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 917 👎 0

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

func connect1(root *Node) *Node {

	// 用层序遍历把每一层的结点全部串起来

	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]
			if i+1 < size {
				queue[i].Next = queue[i+1]

			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	return root
}

func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	var traverse func(nodeA, nodeB *Node)
	traverse = func(nodeA, nodeB *Node) {
		if nil == nodeA || nil == nodeB {
			return
		}

		nodeA.Next = nodeB

		traverse(nodeA.Left, nodeA.Right)
		traverse(nodeB.Left, nodeB.Right)

		// 将整体是做一个三叉树
		// 			+++++++						+++++++
		// 			+nodeA+						+nodeB+
		// 			+++++++						+++++++
		// +++++++			+++++++		+++++++			+++++++
		// + c1  +			+ c2  +		+ c3  +			+ c4  +
		// +++++++			+++++++		+++++++			+++++++
		// [c1]、[c2,c3]、[c4] 视作 [nodeA,nodeB] 的三个子树
		// 在内部将 c2 和 c3 连接起来
		// 为什么可以这样做: 二叉树形态稳定，对于每两层都是固定的结构
		traverse(nodeA.Right, nodeB.Left)
	}

	traverse(root.Left, root.Right)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPopulatingNextRightPointersInEachNode(t *testing.T) {

}
