package leetcode

import (
	"testing"
)

/**
Write an algorithm to determine if a number n is happy.

 A happy number is a number defined by the following process:


 Starting with any positive integer, replace the number by the sum of the
squares of its digits.
 Repeat the process until the number equals 1 (where it will stay), or it loops
endlessly in a cycle which does not include 1.
 Those numbers for which this process ends in 1 are happy.


 Return true if n is a happy number, and false if not.


 Example 1:


Input: n = 19
Output: true
Explanation:
1² + 9² = 82
8² + 2² = 68
6² + 8² = 100
1² + 0² + 0² = 1


 Example 2:


Input: n = 2
Output: false



 Constraints:


 1 <= n <= 2³¹ - 1


 Related Topics 哈希表 数学 双指针 👍 1196 👎 0

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
func isHappy(n int) bool {
	record := make(map[int]bool)

	for !record[n] {
		if n == 1 {
			return true
		}
		record[n] = true
		var newN int
		for ; n != 0; n /= 10 {
			part := n % 10
			newN += part * part
		}
		n = newN
	}

	// 当出现循环且不等于 1 时，表示该数不可能满足条件
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHappyNumber(t *testing.T) {

}
