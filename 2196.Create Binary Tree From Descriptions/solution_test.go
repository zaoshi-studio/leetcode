package leetcode

import (
	"testing"
)

/**
You are given a 2D integer array descriptions where descriptions[i] = [parenti,
childi, isLefti] indicates that parenti is the parent of childi in a binary
tree of unique values. Furthermore,


 If isLefti == 1, then childi is the left child of parenti.
 If isLefti == 0, then childi is the right child of parenti.


 Construct the binary tree described by descriptions and return its root.

 The test cases will be generated such that the binary tree is valid.


 Example 1:


Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
Output: [50,20,80,15,17,19]
Explanation: The root node is the node with value 50 since it has no parent.
The resulting binary tree is shown in the diagram.


 Example 2:


Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
Output: [1,2,null,null,3,4]
Explanation: The root node is the node with value 1 since it has no parent.
The resulting binary tree is shown in the diagram.



 Constraints:


 1 <= descriptions.length <= 10⁴
 descriptions[i].length == 3
 1 <= parenti, childi <= 10⁵
 0 <= isLefti <= 1
 The binary tree described by descriptions is valid.


 Related Topics 树 深度优先搜索 广度优先搜索 数组 哈希表 二叉树 👍 30 👎 0

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
func createBinaryTree(descriptions [][]int) *TreeNode {
	record := make(map[int]*TreeNode, len(descriptions))
	parent := make(map[*TreeNode]*TreeNode)

	for _, v := range descriptions {
		parentVal := v[0]
		childVal := v[1]

		var parentNode, childNode *TreeNode

		if parentNode = record[parentVal]; parentNode == nil {
			parentNode = &TreeNode{
				Val: parentVal,
			}

			record[parentVal] = parentNode
		}

		if childNode = record[childVal]; childNode == nil {
			childNode = &TreeNode{
				Val: childVal,
			}

			record[childVal] = childNode
		}

		parent[childNode] = parentNode

		if v[2] == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}

	for _, v := range record {
		if parent[v] == nil {
			return v
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCreateBinaryTreeFromDescriptions(t *testing.T) {

}
