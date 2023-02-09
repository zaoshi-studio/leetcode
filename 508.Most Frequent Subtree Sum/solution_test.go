package leetcode

import (
	"fmt"
	"testing"
)

/**
Given the root of a binary tree, return the most frequent subtree sum. If there
is a tie, return all the values with the highest frequency in any order.

 The subtree sum of a node is defined as the sum of all the node values formed
by the subtree rooted at that node (including the node itself).


 Example 1:


Input: root = [5,2,-3]
Output: [2,-3,4]


 Example 2:


Input: root = [5,2,-5]
Output: [2]



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 哈希表 二叉树 👍 220 👎 0

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
func findFrequentTreeSum(root *TreeNode) []int {
	var maxCnt int

	record := map[int]int{}

	var postorder func(node *TreeNode) int
	postorder = func(node *TreeNode) int {
		if nil == node {
			return 0
		}

		left := postorder(node.Left)
		right := postorder(node.Right)

		val := left + right + node.Val

		record[val]++

		if record[val] > maxCnt {
			maxCnt = record[val]
		}
		return val
	}

	postorder(root)

	var (
		ans []int
	)

	fmt.Println(record)

	for k, v := range record {
		if v == maxCnt {
			ans = append(ans, k)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMostFrequentSubtreeSum(t *testing.T) {

}
