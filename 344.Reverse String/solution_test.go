package leetcode

import (
	"testing"
)

/**
Write a function that reverses a string. The input string is given as an array
of characters s.

 You must do this by modifying the input array in-place with O(1) extra memory.



 Example 1:
 Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

 Example 2:
 Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]


 Constraints:


 1 <= s.length <= 10⁵
 s[i] is a printable ascii character.


 Related Topics 双指针 字符串 👍 740 👎 0

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
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseString(t *testing.T) {

}
