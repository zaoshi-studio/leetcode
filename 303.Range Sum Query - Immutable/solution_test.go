package leetcode

import (
	"testing"
)

/**
Given an integer array nums, handle multiple queries of the following type:


 Calculate the sum of the elements of nums between indices left and right
inclusive where left <= right.


 Implement the NumArray class:


 NumArray(int[] nums) Initializes the object with the integer array nums.
 int sumRange(int left, int right) Returns the sum of the elements of nums
between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... +
nums[right]).



 Example 1:


Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3



 Constraints:


 1 <= nums.length <= 10⁴
 -10⁵ <= nums[i] <= 10⁵
 0 <= left <= right < nums.length
 At most 10⁴ calls will be made to sumRange.


 Related Topics 设计 数组 前缀和 👍 522 👎 0

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
type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {

	arr := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}

	return NumArray{
		data: arr,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.data[right+1] - this.data[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestRangeSumQueryImmutable(t *testing.T) {

}
