package leetcode

import (
	"testing"
)

/**
Given an integer array nums sorted in non-decreasing order, return an array of
the squares of each number sorted in non-decreasing order.


 Example 1:


Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].


 Example 2:


Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums is sorted in non-decreasing order.



Follow up: Squaring each element and sorting the new array is very trivial,
could you find an
O(n) solution using a different approach?

 Related Topics 数组 双指针 排序 👍 712 👎 0

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

// leetcode submit region begin(Prohibit modification and deletion)
// 原数组按照非递减排序，即使是左侧的负数，经过平方后也可能产生较大的数
// 平方后最大的数，一定是在两端
// 以 0 为分界点，左右两侧是向中间靠拢的趋势
// 相当于用两个指针分别管理左右两侧
// 类似于有序链表的合并
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))

	for i, j, k := 0, len(nums)-1, len(nums)-1; i <= j; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[k] = nums[i] * nums[i]
			i++
		} else {
			ans[k] = nums[j] * nums[j]
			j--
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSquaresOfASortedArray(t *testing.T) {

}
