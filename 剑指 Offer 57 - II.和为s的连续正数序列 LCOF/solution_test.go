package leetcode

import (
	"fmt"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数学 双指针 枚举 👍 497 👎 0

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

func sum(i, j int) int {
	return (i + j) * (j - i + 1) / 2
}

func findContinuousSequence1(target int) [][]int {

	raw := make([]int, target+1)
	for i := 0; i <= target; i++ {
		raw[i] = i
	}

	var ans [][]int

	for i := 1; i <= target; i++ {
		overflow := false
		for j := i + 1; j <= target; j++ {
			if overflow {
				break
			}
			s := sum(i, j)
			if s == target {
				tmp := make([]int, j-i+1)
				copy(tmp, raw[i:j+1])
				ans = append(ans, tmp)
			}

			if s > target {
				overflow = true
			}
		}
	}

	return ans
}

func findContinuousSequence(target int) [][]int {

	var (
		ans [][]int
		s   int
	)

	for left, right := 1, 1; left < target; {
		s = sum(left, right)

		// [left,right] 维护一个窗口
		// 当 s <= target 表示 s 最多满足 target 存在解, 此时向右推进即可
		// 当 s > target 时, 此时一定没有解, 需要收缩窗口
		if s == target {
			var tmp []int
			for k := left; k <= right; k++ {
				tmp = append(tmp, k)
			}

			ans = append(ans, tmp)
			right++
		} else if s < target {
			right++
		} else {
			left++
		}

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHeWeiSdeLianXuZhengShuXuLieLcof(t *testing.T) {
	fmt.Println(findContinuousSequence(9))
}
