package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 栈 设计 队列 👍 640 👎 0

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
type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}

		for len(this.inStack) > 0 {
			//this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			//this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, this.inStack[0])
			this.inStack = this.inStack[1:]
		}
	}

	//v := this.outStack[len(this.outStack)-1]
	//this.outStack = this.outStack[:len(this.outStack)-1]

	v := this.outStack[0]
	this.outStack = this.outStack[1:]
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestYongLiangGeZhanShiXianDuiLieLcof(t *testing.T) {

}
