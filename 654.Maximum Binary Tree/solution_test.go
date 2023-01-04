package leetcode

import (
	"testing"
)

/**
You are given an integer array nums with no duplicates. A maximum binary tree
can be built recursively from nums using the following algorithm:


 Create a root node whose value is the maximum value in nums.
 Recursively build the left subtree on the subarray prefix to the left of the
maximum value.
 Recursively build the right subtree on the subarray suffix to the right of the
maximum value.


 Return the maximum binary tree built from nums.


 Example 1:


Input: nums = [3,2,1,6,0,5]
Output: [6,3,5,null,2,0,null,null,1]
Explanation: The recursive calls are as follow:
- The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right
suffix is [0,5].
    - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix is [
2,1].
        - Empty array, so no child.
        - The largest value in [2,1] is 2. Left prefix is [] and right suffix
is [1].
            - Empty array, so no child.
            - Only one element, so child is a node with value 1.
    - The largest value in [0,5] is 5. Left prefix is [0] and right suffix is []
.
        - Only one element, so child is a node with value 0.
        - Empty array, so no child.


 Example 2:


Input: nums = [3,2,1]
Output: [3,null,2,null,1]



 Constraints:


 1 <= nums.length <= 1000
 0 <= nums[i] <= 1000
 All integers in nums are unique.


 Related Topics 栈 树 数组 分治 二叉树 单调栈 👍 609 👎 0

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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// 以最大值的左右两侧去递归构造二叉树
	var maxIdx int
	for k, v := range nums {
		if v > nums[maxIdx] {
			maxIdx = k
		}
	}

	left := constructMaximumBinaryTree(nums[:maxIdx])
	right := constructMaximumBinaryTree(nums[maxIdx+1:])

	node := &TreeNode{
		Val:   nums[maxIdx],
		Left:  left,
		Right: right,
	}

	return node
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumBinaryTree(t *testing.T) {

}
