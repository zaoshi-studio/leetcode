package leetcode

import (
	"testing"
)

/**
You are playing a solitaire game with three piles of stones of sizes a, b, and
c respectively. Each turn you choose two different non-empty piles, take one
stone from each, and add 1 point to your score. The game stops when there are fewer
than two non-empty piles (meaning there are no more available moves).

 Given three integers a, b, and c, return the maximum score you can get.


 Example 1:


Input: a = 2, b = 4, c = 6
Output: 6
Explanation: The starting state is (2, 4, 6). One optimal set of moves is:
- Take from 1st and 3rd piles, state is now (1, 4, 5)
- Take from 1st and 3rd piles, state is now (0, 4, 4)
- Take from 2nd and 3rd piles, state is now (0, 3, 3)
- Take from 2nd and 3rd piles, state is now (0, 2, 2)
- Take from 2nd and 3rd piles, state is now (0, 1, 1)
- Take from 2nd and 3rd piles, state is now (0, 0, 0)
There are fewer than two non-empty piles, so the game ends. Total: 6 points.


 Example 2:


Input: a = 4, b = 4, c = 6
Output: 7
Explanation: The starting state is (4, 4, 6). One optimal set of moves is:
- Take from 1st and 2nd piles, state is now (3, 3, 6)
- Take from 1st and 3rd piles, state is now (2, 3, 5)
- Take from 1st and 3rd piles, state is now (1, 3, 4)
- Take from 1st and 3rd piles, state is now (0, 3, 3)
- Take from 2nd and 3rd piles, state is now (0, 2, 2)
- Take from 2nd and 3rd piles, state is now (0, 1, 1)
- Take from 2nd and 3rd piles, state is now (0, 0, 0)
There are fewer than two non-empty piles, so the game ends. Total: 7 points.


 Example 3:


Input: a = 1, b = 8, c = 8
Output: 8
Explanation: One optimal set of moves is to take from the 2nd and 3rd piles for
8 turns until they are empty.
After that, there are fewer than two non-empty piles, so the game ends.



 Constraints:


 1 <= a, b, c <= 10⁵


 Related Topics 贪心 数学 堆（优先队列） 👍 52 👎 0

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
func maximumScore(a int, b int, c int) int {
	//if a < 0 || b < 0 || c < 0 {
	//	return 0
	//}
	//
	//if a <= 0 && b <= 0 || a <= 0 && c <= 0 || b <= 0 && c <= 0 {
	//	return 0
	//}
	//
	//if a >= b && a >= c {
	//	return 1 + max(maximumScore(a-1, b-1, c), maximumScore(a-1, b, c-1))
	//} else if b >= a && b >= c {
	//	return 1 + max(maximumScore(a-1, b-1, c), maximumScore(a, b-1, c-1))
	//} else if c >= a && c >= b {
	//	return 1 + max(maximumScore(a-1, b, c-1), maximumScore(a, b-1, c-1))
	//}
	//return 0
	sum := a + b + c
	maxVal := max(max(a, b), c)
	if sum < maxVal*2 {
		return sum - maxVal
	} else {
		return sum / 2
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumScoreFromRemovingStones(t *testing.T) {
	a := 24
	b := 19
	c := 24
	println(maximumScore(a, b, c))
}
