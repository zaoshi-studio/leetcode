package leetcode

import (
	"testing"
)

/**
Given two strings ransomNote and magazine, return true if ransomNote can be
constructed by using the letters from magazine and false otherwise.

 Each letter in magazine can only be used once in ransomNote.


 Example 1:
 Input: ransomNote = "a", magazine = "b"
Output: false

 Example 2:
 Input: ransomNote = "aa", magazine = "ab"
Output: false

 Example 3:
 Input: ransomNote = "aa", magazine = "aab"
Output: true


 Constraints:


 1 <= ransomNote.length, magazine.length <= 10⁵
 ransomNote and magazine consist of lowercase English letters.


 Related Topics 哈希表 字符串 计数 👍 630 👎 0

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
func canConstruct(ransomNote string, magazine string) bool {
	record := make(map[rune]int, len(ransomNote))

	for _, v := range magazine {
		record[v]++
	}

	for _, v := range ransomNote {
		record[v]--
		if record[v] < 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRansomNote(t *testing.T) {

}
