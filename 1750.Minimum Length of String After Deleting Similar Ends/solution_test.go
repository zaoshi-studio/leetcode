package leetcode

import (
	"testing"
)

/**
Given a string s consisting only of characters 'a', 'b', and 'c'. You are asked
to apply the following algorithm on the string any number of times:


 Pick a non-empty prefix from the string s where all the characters in the
prefix are equal.
 Pick a non-empty suffix from the string s where all the characters in this
suffix are equal.
 The prefix and the suffix should not intersect at any index.
 The characters from the prefix and suffix must be the same.
 Delete both the prefix and the suffix.


 Return the minimum length of s after performing the above operation any number
of times (possibly zero times).


 Example 1:


Input: s = "ca"
Output: 2
Explanation: You can't remove any characters, so the string stays as is.


 Example 2:


Input: s = "cabaabac"
Output: 0
Explanation: An optimal sequence of operations is:
- Take prefix = "c" and suffix = "c" and remove them, s = "abaaba".
- Take prefix = "a" and suffix = "a" and remove them, s = "baab".
- Take prefix = "b" and suffix = "b" and remove them, s = "aa".
- Take prefix = "a" and suffix = "a" and remove them, s = "".

 Example 3:


Input: s = "aabccabba"
Output: 3
Explanation: An optimal sequence of operations is:
- Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
- Take prefix = "b" and suffix = "bb" and remove them, s = "cca".



 Constraints:


 1 <= s.length <= 10⁵
 s only consists of characters 'a', 'b', and 'c'.


 Related Topics 双指针 字符串 👍 39 👎 0

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
func minimumLength(s string) int {

	if len(s) == 0 {
		return 0
	}

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			break
		}

		common := s[i]

		for i < j && s[i] == common {
			i++
		}

		if i == j {
			return 0
		}

		for i < j && s[j] == common {
			j--
		}
	}

	return j - i + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumLengthOfStringAfterDeletingSimilarEnds(t *testing.T) {

}
