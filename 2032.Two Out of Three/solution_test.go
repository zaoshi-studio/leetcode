package leetcode

import (
	"testing"
)

/**
Given three integer arrays nums1, nums2, and nums3, return a distinct array
containing all the values that are present in at least two out of the three arrays.
You may return the values in any order.


 Example 1:


Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
Output: [3,2]
Explanation: The values that are present in at least two arrays are:
- 3, in all three arrays.
- 2, in nums1 and nums2.


 Example 2:


Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
Output: [2,3,1]
Explanation: The values that are present in at least two arrays are:
- 2, in nums2 and nums3.
- 3, in nums1 and nums2.
- 1, in nums1 and nums3.


 Example 3:


Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
Output: []
Explanation: No value is present in at least two arrays.



 Constraints:


 1 <= nums1.length, nums2.length, nums3.length <= 100
 1 <= nums1[i], nums2[j], nums3[k] <= 100


 Related Topics 数组 哈希表 👍 44 👎 0

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
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {

	var ans []int

	data := make(map[int]int, len(nums1)+len(nums2)+len(nums3))

	for _, v := range nums1 {
		data[v] |= 1 << 1
	}

	for _, v := range nums2 {
		data[v] |= 1 << 2
	}

	for _, v := range nums3 {
		data[v] |= 1 << 3
	}

	for k, v := range data {
		//if v == 3 || v == 5 || v == 6 || v == 7 {
		//	ans = append(ans, k)
		//}

		if v&(v-1) > 0 {
			ans = append(ans, k)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoOutOfThree(t *testing.T) {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}
	twoOutOfThree(nums1, nums2, nums3)
}
