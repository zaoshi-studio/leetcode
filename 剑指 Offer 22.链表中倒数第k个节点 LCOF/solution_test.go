package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 链表 双指针 👍 419 👎 0

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

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {

	//dummy := &ListNode{Next: head}

	p1 := head
	cnt := 0

	for ; cnt != k; p1 = p1.Next {
		cnt++
	}

	p2 := head
	for ; p1 != nil; p1 = p1.Next {
		p2 = p2.Next
	}
	return p2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLianBiaoZhongDaoShuDiKgeJieDianLcof(t *testing.T) {

}
