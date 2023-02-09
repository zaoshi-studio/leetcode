package leetcode

import (
	"testing"
)

/**
Implement the BSTIterator class that represents an iterator over the in-order
traversal of a binary search tree (BST):


 BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The
root of the BST is given as part of the constructor. The pointer should be
initialized to a non-existent number smaller than any element in the BST.
 boolean hasNext() Returns true if there exists a number in the traversal to
the right of the pointer, otherwise returns false.
 int next() Moves the pointer to the right, then returns the number at the
pointer.


 Notice that by initializing the pointer to a non-existent smallest number, the
first call to next() will return the smallest element in the BST.

 You may assume that next() calls will always be valid. That is, there will be
at least a next number in the in-order traversal when next() is called.


 Example 1:


Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext",
 "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]


Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next(); // return 3
bSTIterator.next(); // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 20
bSTIterator.hasNext(); // return False



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁵].
 0 <= Node.val <= 10⁶
 At most 10⁵ calls will be made to hasNext, and next.



 Follow up:


 Could you implement next() and hasNext() to run in average O(1) time and use O(
h) memory, where h is the height of the tree?


 Related Topics 栈 树 设计 二叉搜索树 二叉树 迭代器 👍 662 👎 0

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
type BSTIterator struct {
	data []int
	idx  int
}

func Constructor(root *TreeNode) BSTIterator {

	iterator := BSTIterator{
		data: []int{},
	}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if nil == node {
			return
		}

		preorder(node.Left)

		iterator.data = append(iterator.data, node.Val)

		preorder(node.Right)
	}

	preorder(root)

	return iterator
}

func (this *BSTIterator) Next() int {
	v := this.data[this.idx]

	this.idx++

	return v
}

func (this *BSTIterator) HasNext() bool {
	return this.idx != len(this.data)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestBinarySearchTreeIterator(t *testing.T) {

}
