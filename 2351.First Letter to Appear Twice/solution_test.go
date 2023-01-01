package leetcode

import (
	"testing"
)

/**
Given a string s consisting of lowercase English letters, return the first
letter to appear twice.

 Note:


 A letter a appears twice before another letter b if the second occurrence of a
is before the second occurrence of b.
 s will contain at least one letter that appears twice.



 Example 1:


Input: s = "abccbaacz"
Output: "c"
Explanation:
The letter 'a' appears on the indexes 0, 5 and 6.
The letter 'b' appears on the indexes 1 and 4.
The letter 'c' appears on the indexes 2, 3 and 7.
The letter 'z' appears on the index 8.
The letter 'c' is the first letter to appear twice, because out of all the
letters the index of its second occurrence is the smallest.


 Example 2:


Input: s = "abcdd"
Output: "d"
Explanation:
The only letter that appears twice is 'd' so we return 'd'.



 Constraints:


 2 <= s.length <= 100
 s consists of lowercase English letters.
 s has at least one repeated letter.


 Related Topics 哈希表 字符串 计数 👍 31 👎 0

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
func repeatedCharacter1(s string) byte {
	//record := make(map[byte]int, len(s))
	//
	//for i := 0; i < len(s); i++ {
	//	record[s[i]]++
	//	if record[s[i]] >= 2 {
	//		return s[i]
	//	}
	//}

	bit := 0

	for _, v := range s {
		n := 1 << (v - 'a')

		if bit&n == n {
			return byte(v)
		}
		bit |= n
	}

	return 0
}

func repeatedCharacter(s string) byte {

	bit := 0

	for _, v := range s {
		n := 1 << (v - 'a')

		if bit&n == n {
			return byte(v)
		}
		bit |= n
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstLetterToAppearTwice(t *testing.T) {
	s := "abccbaacz"
	repeatedCharacter(s)
}
