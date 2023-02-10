package leetcode

import (
	"testing"
)

/**
Given a binary tree with the following rules:


 root.val == 0
 If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 *
x + 1
 If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 *
 x + 2


 Now the binary tree is contaminated, which means all treeNode.val have been
changed to -1.

 Implement the FindElements class:


 FindElements(TreeNode* root) Initializes the object with a contaminated binary
tree and recovers it.
 bool find(int target) Returns true if the target value exists in the recovered
binary tree.



 Example 1:


Input
["FindElements","find","find"]
[[[-1,null,-1]],[1],[2]]
Output
[null,false,true]
Explanation
FindElements findElements = new FindElements([-1,null,-1]);
findElements.find(1); // return False
findElements.find(2); // return True

 Example 2:


Input
["FindElements","find","find","find"]
[[[-1,-1,-1,-1,-1]],[1],[3],[5]]
Output
[null,true,true,false]
Explanation
FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
findElements.find(1); // return True
findElements.find(3); // return True
findElements.find(5); // return False

 Example 3:


Input
["FindElements","find","find","find","find"]
[[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
Output
[null,true,false,false,true]
Explanation
FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
findElements.find(2); // return True
findElements.find(3); // return False
findElements.find(4); // return False
findElements.find(5); // return True



 Constraints:


 TreeNode.val == -1
 The height of the binary tree is less than or equal to 20
 The total number of nodes is between [1, 10⁴]
 Total calls of find() is between [1, 10⁴]
 0 <= target <= 10⁶


 Related Topics 树 深度优先搜索 广度优先搜索 设计 哈希表 二叉树 👍 41 👎 0

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
type FindElements struct {
	record map[int]bool
}

func Constructor(root *TreeNode) FindElements {

	elem := FindElements{
		record: map[int]bool{},
	}

	if root != nil {
		root.Val = 0
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}

		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
		}

		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
		}

		elem.record[node.Val] = true

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return elem
}

func (this *FindElements) Find(target int) bool {
	return this.record[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestFindElementsInAContaminatedBinaryTree(t *testing.T) {

}
