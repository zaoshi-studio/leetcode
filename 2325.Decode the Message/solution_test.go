package leetcode

import (
	"strings"
	"testing"
)

/**
You are given the strings key and message, which represent a cipher key and a
secret message, respectively. The steps to decode message are as follows:


 Use the first appearance of all 26 lowercase English letters in key as the
order of the substitution table.
 Align the substitution table with the regular English alphabet.
 Each letter in message is then substituted using the table.
 Spaces ' ' are transformed to themselves.



 For example, given key = "happy boy" (actual key would have at least one
instance of each letter in the alphabet), we have the partial substitution table of (
'h' -> 'a', 'a' -> 'b', 'p' -> 'c', 'y' -> 'd', 'b' -> 'e', 'o' -> 'f').


 Return the decoded message.


 Example 1:


Input: key = "the quick brown fox jumps over the lazy dog", message = "vkbs bs
t suepuv"
Output: "this is a secret"
Explanation: The diagram above shows the substitution table.
It is obtained by taking the first appearance of each letter in "the quick
brown fox jumps over the lazy dog".


 Example 2:


Input: key = "eljuxhpwnyrdgtqkviszcfmabo", message = "zwx hnfx lqantp mnoeius
ycgk vcnjrdb"
Output: "the five boxing wizards jump quickly"
Explanation: The diagram above shows the substitution table.
It is obtained by taking the first appearance of each letter in
"eljuxhpwnyrdgtqkviszcfmabo".



 Constraints:


 26 <= key.length <= 2000
 key consists of lowercase English letters and ' '.
 key contains every letter in the English alphabet ('a' to 'z') at least once.
 1 <= message.length <= 2000
 message consists of lowercase English letters and ' '.


 Related Topics 哈希表 字符串 👍 34 👎 0

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
func decodeMessage(key string, message string) string {
	builder := strings.Builder{}

	var idx byte

	hashMap := make(map[byte]byte, 26)

	for i := 0; i < len(key); i++ {
		if key[i] == ' ' {
			continue
		}
		// the current character key[i] is already used
		if _, ok := hashMap[key[i]]; ok {
			continue
		}

		hashMap[key[i]] = 'a' + idx
		idx++
	}

	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			builder.WriteByte(' ')
			continue
		}
		builder.WriteByte(hashMap[message[i]])
	}

	return builder.String()
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDecodeTheMessage(t *testing.T) {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	println(decodeMessage(key, message))
}
