package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 栈 树 设计 二叉搜索树 二叉树 迭代器 👍 39 👎 0

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
	ptr  int
}

func Constructor(root *TreeNode) BSTIterator {

	iterator := BSTIterator{
		data: []int{},
		ptr:  0,
	}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
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
	val := this.data[this.ptr]
	this.ptr++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.ptr < len(this.data)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestKTOapQ(t *testing.T) {

}
