package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 哈希表 链表 👍 629 👎 0

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

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	dummy := &Node{}
	cur := dummy

	old2New := map[*Node]*Node{}

	for p := head; p != nil; {

		node := &Node{
			Val: p.Val,
		}

		old2New[p] = node

		if rNode, ok := old2New[p.Random]; ok {
			node.Random = rNode
		}

		cur.Next = node
		cur = node

		for k, v := range old2New {
			if k.Random == p {
				v.Random = node
			}
		}

		p = p.Next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFuZaLianBiaoDeFuZhiLcof(t *testing.T) {

}
