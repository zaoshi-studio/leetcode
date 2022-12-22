package leetcode

import (
	"testing"
)

/**
You are given an integer array height of length n. There are n vertical lines
drawn such that the two endpoints of the iᵗʰ line are (i, 0) and (i, height[i]).

 Find two lines that together with the x-axis form a container, such that the
container contains the most water.

 Return the maximum amount of water a container can store.

 Notice that you may not slant the container.


 Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,
7]. In this case, the max area of water (blue section) the container can
contain is 49.


 Example 2:


Input: height = [1,1]
Output: 1



 Constraints:


 n == height.length
 2 <= n <= 10⁵
 0 <= height[i] <= 10⁴


 Related Topics 贪心 数组 双指针 👍 3976 👎 0

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
func maxArea(height []int) int {

	area := 0

	for left, right := 0, len(height)-1; left < right; {

		// 计算底部长度
		l := right - left
		// 计算高度
		h := -max(-height[left], -height[right])
		// 计算容积
		area = max(l*h, area)
		// 收缩容器
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {

}
