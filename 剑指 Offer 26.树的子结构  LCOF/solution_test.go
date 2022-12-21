package leetcode

import(
    "testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 深度优先搜索 二叉树 👍 670 👎 0

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct{
    Val int
    Left *TreeNode
    Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		// 当前根节点值不相等, 尝试在子节点中找到能否与 B 匹配的部分
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	} else {
		// 根节点相等, 检测子节点
		// 虽然结点值相等, 但要考虑有多个节点值相等
		// 不能在第一次相等,但子节点不匹配时直接返回 false
		return compare(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}

}

// compare 标记两个节点的根节点已经相等, 进而比较子节点
func compare(nodeA, nodeB *TreeNode) bool {
	if nodeB == nil {
		return true
	}

	if nodeA == nil || nodeA.Val != nodeB.Val {
		return false
	}

	return compare(nodeA.Left, nodeB.Left) && compare(nodeA.Right, nodeB.Right)
}
//leetcode submit region end(Prohibit modification and deletion)


func TestShuDeZiJieGouLcof(t *testing.T){
    
}