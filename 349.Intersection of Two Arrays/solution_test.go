package leetcode

import (
	"sort"
	"testing"
)

/**
Given two integer arrays nums1 and nums2, return an array of their intersection.
 Each element in the result must be unique and you may return the result in any
order.


 Example 1:


Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]


 Example 2:


Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.



 Constraints:


 1 <= nums1.length, nums2.length <= 1000
 0 <= nums1[i], nums2[i] <= 1000


 Related Topics 数组 哈希表 双指针 二分查找 排序 👍 711 👎 0

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
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var ans []int

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		v1, v2 := nums1[i], nums2[j]

		if v1 == v2 {
			if len(ans) == 0 || ans[len(ans)-1] != v1 {
				ans = append(ans, v1)
			}
			i++
			j++
		} else if v1 > v2 {
			j++
		} else {
			i++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArrays(t *testing.T) {

}
