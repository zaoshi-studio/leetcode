package leetcode

import (
	"math"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 栈 设计 👍 417 👎 0

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
type MinStack struct {
	data     []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:     []int{},
		minStack: []int{math.MaxInt},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	m := min(x, this.minStack[len(this.minStack)-1])
	this.minStack = append(this.minStack, m)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestBaoHanMinhanShuDeZhanLcof(t *testing.T) {

}
