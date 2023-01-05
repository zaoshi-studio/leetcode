package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 字符串 回溯 👍 630 👎 0

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
func permutation(s string) []string {

	raw := []byte(s)

	sort.Slice(raw, func(i, j int) bool {
		return raw[i] < raw[j]
	})

	var ans []string

	record := make([]bool, len(s))

	var backtrack func(idx int, buffer []byte)
	backtrack = func(idx int, buffer []byte) {
		if len(buffer) == len(s) {
			tmp := make([]byte, len(s))
			copy(tmp, buffer)
			ans = append(ans, string(tmp))
			return
		}

		for i := 0; i < len(s); i++ {

			if record[i] || i > 0 && raw[i] == raw[i-1] && !record[i-1] {
				continue
			}
			record[i] = true
			buffer = append(buffer, raw[i])
			backtrack(i, buffer)
			buffer = buffer[:len(buffer)-1]
			record[i] = false
		}
	}

	backtrack(0, []byte{})

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

func TestZiFuChuanDePaiLieLcof(t *testing.T) {

}
