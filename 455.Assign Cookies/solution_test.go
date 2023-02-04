package leetcode

import (
	"sort"
	"testing"
)

/**
Assume you are an awesome parent and want to give your children some cookies.
But, you should give each child at most one cookie.

 Each child i has a greed factor g[i], which is the minimum size of a cookie
that the child will be content with; and each cookie j has a size s[j]. If s[j] >=
g[i], we can assign the cookie j to the child i, and the child i will be
content. Your goal is to maximize the number of your content children and output the
maximum number.


 Example 1:


Input: g = [1,2,3], s = [1,1]
Output: 1
Explanation: You have 3 children and 2 cookies. The greed factors of 3 children
are 1, 2, 3.
And even though you have 2 cookies, since their size is both 1, you could only
make the child whose greed factor is 1 content.
You need to output 1.


 Example 2:


Input: g = [1,2], s = [1,2,3]
Output: 2
Explanation: You have 2 children and 3 cookies. The greed factors of 2 children
are 1, 2.
You have 3 cookies and their sizes are big enough to gratify all of the
children,
You need to output 2.



 Constraints:


 1 <= g.length <= 3 * 10⁴
 0 <= s.length <= 3 * 10⁴
 1 <= g[i], s[j] <= 2³¹ - 1


 Related Topics 贪心 数组 双指针 排序 👍 640 👎 0

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
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var ans int

	// 优先喂最小的
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			ans++
			i++
		}
	}

	// 优先喂最大的
	// 只能控制 胃口 g[i] 递减
	// 当 g[i] > s[j] 时只能减少 g[i] 才可能去找到满足的条件
	// 如果控制 饼干 s[j] 递减
	// 当 g[i] > s[j] 时操作起来相对麻烦，提高不了饼干的数值
	//for i, j := len(g)-1, len(s)-1; i >= 0; i-- {
	//	if j >= 0 && s[j] >= g[i] {
	//		ans++
	//		j--
	//	}
	//}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAssignCookies(t *testing.T) {

}
