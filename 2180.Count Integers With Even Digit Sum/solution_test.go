package leetcode

import (
	"testing"
)

/**
Given a positive integer num, return the number of positive integers less than
or equal to num whose digit sums are even.

 The digit sum of a positive integer is the sum of all its digits.


 Example 1:


Input: num = 4
Output: 2
Explanation:
The only integers less than or equal to 4 whose digit sums are even are 2 and 4.



 Example 2:


Input: num = 30
Output: 14
Explanation:
The 14 integers less than or equal to 30 whose digit sums are even are
2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.



 Constraints:


 1 <= num <= 1000


 Related Topics 数学 模拟 👍 47 👎 0

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
func countEven(num int) int {
	var ans int

	for i := 1; i <= num; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}

		if sum%2 == 0 {
			ans++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountIntegersWithEvenDigitSum(t *testing.T) {

}
