package leetcode

import (
	"testing"
)

/**
Given a binary tree, design an algorithm which creates a linked list of all the
nodes at each depth (e.g., if you have a tree with depth D, you'll have D
linked lists). Return a array containing all the linked lists.



 Example:


Input: [1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

Output: [[1],[2,3],[4,5,7],[8]]


 Related Topics 树 广度优先搜索 链表 二叉树 👍 88 👎 0

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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	var ans []*ListNode

	queue := []*TreeNode{tree}

	for len(queue) != 0 {
		size := len(queue)

		head := &ListNode{}
		tail := head

		for i := 0; i < size; i++ {
			node := queue[i]

			listNode := &ListNode{
				Val: node.Val,
			}

			tail.Next = listNode
			tail = listNode

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]

		ans = append(ans, head.Next)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestListOfDepthLcci(t *testing.T) {

}
