package leetcode

import (
	"sort"
	"testing"
)

/**
Given an array nums of n integers, return an array of all the unique
quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:


 0 <= a, b, c, d < n
 a, b, c, and d are distinct.
 nums[a] + nums[b] + nums[c] + nums[d] == target


 You may return the answer in any order.


 Example 1:


Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]


 Example 2:


Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]



 Constraints:


 1 <= nums.length <= 200
 -10⁹ <= nums[i] <= 10⁹
 -10⁹ <= target <= 10⁹


 Related Topics 数组 双指针 排序 👍 1486 👎 0

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
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var ans [][]int

	for i := 0; i < len(nums)-3; i++ {
		// 确定第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 确定第二个数
		for j := i + 1; j < len(nums)-2; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 采用双指针确定最后两个数

			twoSum := nums[i] + nums[j]

			for left, right := j+1, len(nums)-1; left < right; {
				sum := twoSum + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})

					// 去重
					// 此处 left 向前比，right 向后比
					// 当比较完成时，left 和 right 都会落在与原先值相等的位置，需要再推进一步
					for ; left < right && nums[left] == nums[left+1]; left++ {
					}
					left++

					for ; left < right && nums[right] == nums[right-1]; right-- {
					}
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFourSum(t *testing.T) {

}
