package leetcode

import (
	"fmt"
	"testing"
)

/**
A complete binary tree is a binary tree in which every level, except possibly
the last, is completely filled, and all nodes are as far left as possible.

 Design an algorithm to insert a new node to a complete binary tree keeping it
complete after the insertion.

 Implement the CBTInserter class:


 CBTInserter(TreeNode root) Initializes the data structure with the root of the
complete binary tree.
 int insert(int v) Inserts a TreeNode into the tree with value Node.val == val
so that the tree remains complete, and returns the value of the parent of the
inserted TreeNode.
 TreeNode get_root() Returns the root node of the tree.



 Example 1:


Input
["CBTInserter", "insert", "insert", "get_root"]
[[[1, 2]], [3], [4], []]
Output
[null, 1, 2, [1, 2, 3, 4]]


Explanation
CBTInserter cBTInserter = new CBTInserter([1, 2]);
cBTInserter.insert(3); // return 1
cBTInserter.insert(4); // return 2
cBTInserter.get_root(); // return [1, 2, 3, 4]



 Constraints:


 The number of nodes in the tree will be in the range [1, 1000].
 0 <= Node.val <= 5000
 root is a complete binary tree.
 0 <= val <= 5000
 At most 10⁴ calls will be made to insert and get_root.


 Related Topics 树 广度优先搜索 设计 二叉树 👍 148 👎 0

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
type CBTInserter struct {
	root        *TreeNode
	unfullQueue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {

	fmt.Println("constructor")

	ret := CBTInserter{
		root:        root,
		unfullQueue: []*TreeNode{},
	}

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

			if node.Left == nil || node.Right == nil {
				ret.unfullQueue = append(ret.unfullQueue, node)
			}
		}

		queue = queue[size:]
	}

	for _, v := range ret.unfullQueue {
		fmt.Println(v.Val)
	}

	return ret
}

func (this *CBTInserter) Insert(val int) int {

	fmt.Println("insert")

	newNode := &TreeNode{
		Val: val,
	}

	pNode := this.unfullQueue[0]

	if pNode.Left == nil {
		pNode.Left = newNode
	} else if pNode.Right == nil {
		pNode.Right = newNode
		this.unfullQueue = this.unfullQueue[1:]
	}

	this.unfullQueue = append(this.unfullQueue, newNode)

	return pNode.Val
}

func (this *CBTInserter) Get_root() *TreeNode {

	fmt.Println("get root")

	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestCompleteBinaryTreeInserter(t *testing.T) {

}
