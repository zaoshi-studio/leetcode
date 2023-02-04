package leetcode

import (
	"testing"
)

/**
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n,
return the number of tuples (i, j, k, l) such that:


 0 <= i, j, k, l < n
 nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0



 Example 1:


Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) +
 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) +
 0 = 0


 Example 2:


Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
Output: 1



 Constraints:


 n == nums1.length
 n == nums2.length
 n == nums3.length
 n == nums4.length
 1 <= n <= 200
 -2²⁸ <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2²⁸


 Related Topics 数组 哈希表 👍 733 👎 0

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
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 题目中，每个数组都必须取一个元素
	// 相当于四个数组做笛卡尔积，只是采用加法的形式
	// 为了计算简单 两个一组，对两个数组做笛卡尔积更为方便，同时降低复杂度
	n := len(nums1)
	var ans int
	part1 := make(map[int]int, n)

	for _, v := range nums1 {
		for _, vv := range nums2 {
			part1[v+vv]++
		}
	}

	for _, v := range nums3 {
		for _, vv := range nums4 {
			ans += part1[-v-vv]
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFourSumIi(t *testing.T) {

}
