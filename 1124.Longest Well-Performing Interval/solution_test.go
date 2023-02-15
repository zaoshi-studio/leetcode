package leetcode

import (
	"testing"
)

/**
We are given hours, a list of the number of hours worked per day for a given
employee.

 A day is considered to be a tiring day if and only if the number of hours
worked is (strictly) greater than 8.

 A well-performing interval is an interval of days for which the number of
tiring days is strictly larger than the number of non-tiring days.

 Return the length of the longest well-performing interval.


 Example 1:


Input: hours = [9,9,6,0,6,6,9]
Output: 3
Explanation: The longest well-performing interval is [9,9,6].


 Example 2:


Input: hours = [6,6,6]
Output: 0



 Constraints:


 1 <= hours.length <= 10⁴
 0 <= hours[i] <= 16


 Related Topics 栈 数组 哈希表 前缀和 单调栈 👍 337 👎 0

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
func longestWPI(hours []int) (ans int) {

	// 连续表现良好时间段 => 前缀和 s > 0
	s := 0
	pos := map[int]int{}
	for i, x := range hours {
		if x > 8 {
			s++
		} else {
			s--
		}
		if s > 0 {
			// 只要当前前缀和满足条件，则之前的所有天数都能被算入连续的时间段
			ans = i + 1
		} else if j, ok := pos[s-1]; ok {
			// 此时 s == 0 或者 s < 0
			// j 记录的是 低谷 的起始位置
			// 防止 <8 的部分连续出现，导致 >8 的部分无法被算进去
			// 例子 6,6,9
			ans = max(ans, i-j)
		}
		if _, ok := pos[s]; !ok {
			pos[s] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func longestWPI(hours []int) int {
//	var (
//		ans int
//		cnt int
//	)
//
//	for i, j := 0, 0; j < len(hours); j++ {
//		if cnt == 0 && hours[j] > 8 {
//			cnt++
//			for cnt > 0 {
//
//			}
//		}
//	}
//
//	return ans
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestWellPerformingInterval(t *testing.T) {
	//hours := []int{9, 9, 6, 0, 6, 6, 9}
	//hours := []int{9, 9, 6, 0, 9, 9, 9, 9}
	hours := []int{6, 6, 9}
	longestWPI(hours)
}
