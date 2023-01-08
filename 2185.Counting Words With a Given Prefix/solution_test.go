package leetcode

import (
	"regexp"
	"strings"
	"testing"
)

/**
You are given an array of strings words and a string pref.

 Return the number of strings in words that contain pref as a prefix.

 A prefix of a string s is any leading contiguous substring of s.


 Example 1:


Input: words = ["pay","attention","practice","attend"], pref = "at"
Output: 2
Explanation: The 2 strings that contain "at" as a prefix are: "attention" and
"attend".


 Example 2:


Input: words = ["leetcode","win","loops","success"], pref = "code"
Output: 0
Explanation: There are no strings that contain "code" as a prefix.



 Constraints:


 1 <= words.length <= 100
 1 <= words[i].length, pref.length <= 100
 words[i] and pref consist of lowercase English letters.


 Related Topics 数组 字符串 👍 25 👎 0

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
func prefixCount1(words []string, pref string) int {
	var ans int

	reg, _ := regexp.Compile("^" + pref)

	for _, v := range words {
		if reg.MatchString(v) {
			ans++
		}
	}

	return ans
}

func prefixCount(words []string, pref string) int {
	var ans int

	for _, v := range words {
		if strings.HasPrefix(v, pref) {
			ans++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountingWordsWithAGivenPrefix(t *testing.T) {

}
