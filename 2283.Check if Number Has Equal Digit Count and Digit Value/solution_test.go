package leetcode

import (
	"testing"
)

/**
You are given a 0-indexed string num of length n consisting of digits.

 Return true if for every index i in the range 0 <= i < n, the digit i occurs
num[i] times in num, otherwise return false.


 Example 1:


Input: num = "1210"
Output: true
Explanation:
num[0] = '1'. The digit 0 occurs once in num.
num[1] = '2'. The digit 1 occurs twice in num.
num[2] = '1'. The digit 2 occurs once in num.
num[3] = '0'. The digit 3 occurs zero times in num.
The condition holds true for every index in "1210", so return true.


 Example 2:


Input: num = "030"
Output: false
Explanation:
num[0] = '0'. The digit 0 should occur zero times, but actually occurs twice in
num.
num[1] = '3'. The digit 1 should occur three times, but actually occurs zero
times in num.
num[2] = '0'. The digit 2 occurs zero times in num.
The indices 0 and 1 both violate the condition, so return false.



 Constraints:


 n == num.length
 1 <= n <= 10
 num consists of digits.


 Related Topics 哈希表 字符串 计数 👍 30 👎 0

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
func digitCount(num string) bool {
	record := make(map[int]int, len(num))

	for i := 0; i < len(num); i++ {
		record[int(num[i])-48]++
	}

	for i := 0; i < len(num); i++ {
		idx := int(num[i]) - 48
		if record[i] != idx {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCheckIfNumberHasEqualDigitCountAndDigitValue(t *testing.T) {
	num := "1210"
	digitCount(num)
}
