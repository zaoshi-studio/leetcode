package leetcode

import (
	"sort"
	"testing"
)

/**
You have a water dispenser that can dispense cold, warm, and hot water. Every
second, you can either fill up 2 cups with different types of water, or 1 cup of
any type of water.

 You are given a 0-indexed integer array amount of length 3 where amount[0],
amount[1], and amount[2] denote the number of cold, warm, and hot water cups you
need to fill respectively. Return the minimum number of seconds needed to fill up
all the cups.


 Example 1:


Input: amount = [1,4,2]
Output: 4
Explanation: One way to fill up the cups is:
Second 1: Fill up a cold cup and a warm cup.
Second 2: Fill up a warm cup and a hot cup.
Second 3: Fill up a warm cup and a hot cup.
Second 4: Fill up a warm cup.
It can be proven that 4 is the minimum number of seconds needed.


 Example 2:


Input: amount = [5,4,4]
Output: 7
Explanation: One way to fill up the cups is:
Second 1: Fill up a cold cup, and a hot cup.
Second 2: Fill up a cold cup, and a warm cup.
Second 3: Fill up a cold cup, and a warm cup.
Second 4: Fill up a warm cup, and a hot cup.
Second 5: Fill up a cold cup, and a hot cup.
Second 6: Fill up a cold cup, and a warm cup.
Second 7: Fill up a hot cup.


 Example 3:


Input: amount = [5,0,0]
Output: 5
Explanation: Every second, we fill up a cold cup.



 Constraints:


 amount.length == 3
 0 <= amount[i] <= 100


 Related Topics 贪心 数组 排序 堆（优先队列） 👍 60 👎 0

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
func fillCups1(amount []int) int {
	var ans int
	for {
		sort.Ints(amount)
		if amount[0] > 0 {
			amount[0]--
			amount[2]--
			ans++
		} else if amount[1] > 0 {
			amount[1]--
			amount[2]--
			ans++
		} else {
			ans += amount[2]
			break
		}
	}

	return ans
}

func fillCups(amount []int) int {
	sort.Ints(amount)
	if amount[2] > amount[1]+amount[0] {
		return amount[2]
	}

	// +1 为了向上取整
	return (amount[2] + amount[1] + amount[0] + 1) / 2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumAmountOfTimeToFillCups(t *testing.T) {

}
