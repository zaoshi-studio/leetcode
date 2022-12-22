package leetcode

import (
	"math"
	"sort"
	"testing"
)

/**
Given an integer array nums of length n and an integer target, find three
integers in nums such that the sum is closest to target.

 Return the sum of the three integers.

 You may assume that each input would have exactly one solution.


 Example 1:


Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).


 Example 2:


Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).



 Constraints:


 3 <= nums.length <= 500
 -1000 <= nums[i] <= 1000
 -10⁴ <= target <= 10⁴


 Related Topics 数组 双指针 排序 👍 1304 👎 0

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
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	ans := math.MaxInt

	for i := 0; i < len(nums); i++ {
		sum := 0

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {

			sum = nums[i] + nums[j] + nums[k]

			if sum == target {
				return target
			}

			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}

		}
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSumClosest(t *testing.T) {

}
