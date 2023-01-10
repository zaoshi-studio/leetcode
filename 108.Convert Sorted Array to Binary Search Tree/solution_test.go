package leetcode

import (
	"testing"
)

/**
Given an integer array nums where the elements are sorted in ascending order,
convert it to a height-balanced binary search tree.


 Example 1:


Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:



 Example 2:


Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums is sorted in a strictly increasing order.


 Related Topics 树 二叉搜索树 数组 分治 二叉树 👍 1214 👎 0

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
func sortedArrayToBST(nums []int) *TreeNode {

	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {

		if left < 0 || left > right || right > len(nums)-1 {
			return nil
		}

		// 选取中间节点作为根节点
		mid := (left + right) / 2
		root := &TreeNode{
			Val: nums[mid],
		}

		root.Left = build(left, mid-1)
		root.Right = build(mid+1, right)

		return root
	}

	return build(0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertSortedArrayToBinarySearchTree(t *testing.T) {

}
