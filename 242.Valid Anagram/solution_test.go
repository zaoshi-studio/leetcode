package leetcode

import (
	"testing"
)

/**
Given two strings s and t, return true if t is an anagram of s, and false
otherwise.

 An Anagram is a word or phrase formed by rearranging the letters of a
different word or phrase, typically using all the original letters exactly once.


 Example 1:
 Input: s = "anagram", t = "nagaram"
Output: true

 Example 2:
 Input: s = "rat", t = "car"
Output: false


 Constraints:


 1 <= s.length, t.length <= 5 * 10⁴
 s and t consist of lowercase English letters.



 Follow up: What if the inputs contain Unicode characters? How would you adapt
your solution to such a case?

 Related Topics 哈希表 字符串 排序 👍 724 👎 0

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
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[rune]int, len(s))

	for _, v := range s {
		hashMap[v]++
	}

	for _, v := range t {
		hashMap[v]--
		if hashMap[v] < 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {

}
