package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 505 👎 0

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
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)

	return arr[:k]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestZuiXiaoDeKgeShuLcof(t *testing.T) {
	arr := []int{3, 2, 1}
	k := 2
	getLeastNumbers(arr, k)
}
