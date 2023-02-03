package leetcode

import (
	"testing"
)

/**
Design your implementation of the linked list. You can choose to use a singly
or doubly linked list. A node in a singly linked list should have two attributes:
val and next. val is the value of the current node, and next is a pointer/
reference to the next node. If you want to use the doubly linked list, you will need
one more attribute prev to indicate the previous node in the linked list. Assume
all nodes in the linked list are 0-indexed.

 Implement the MyLinkedList class:


 MyLinkedList() Initializes the MyLinkedList object.
 int get(int index) Get the value of the indexᵗʰ node in the linked list. If
the index is invalid, return -1.
 void addAtHead(int val) Add a node of value val before the first element of
the linked list. After the insertion, the new node will be the first node of the
linked list.
 void addAtTail(int val) Append a node of value val as the last element of the
linked list.
 void addAtIndex(int index, int val) Add a node of value val before the indexᵗʰ
node in the linked list. If index equals the length of the linked list, the
node will be appended to the end of the linked list. If index is greater than the
length, the node will not be inserted.
 void deleteAtIndex(int index) Delete the indexᵗʰ node in the linked list, if
the index is valid.



 Example 1:


Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex",
 "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3



 Constraints:


 0 <= index, val <= 1000
 Please do not use the built-in LinkedList library.
 At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and
deleteAtIndex.


 Related Topics 设计 链表 👍 735 👎 0

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
type MyLinkedList struct {
	head, tail *MyLinkedListNode
	length     int
}

type MyLinkedListNode struct {
	next *MyLinkedListNode
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head:   &MyLinkedListNode{},
		tail:   &MyLinkedListNode{},
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	ans := -1

	for p, cnt := this.head, -1; p != nil; p = p.next {
		if cnt == index {
			ans = p.val
			break
		}
		cnt++
	}


	return ans
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {

	this.AddAtIndex(this.length, val)

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &MyLinkedListNode{
		val: val,
	}

	for p, cnt := this.head, -1; p != nil; p = p.next {
		if cnt+1 == index {
			next := p.next
			p.next = newNode
			newNode.next = next
			break
		}
		cnt++
	}

	if this.length == 0 {
		this.tail = newNode
	}

	if this.length+1 == index {
		this.tail = newNode
	}
	this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}

	var prev *MyLinkedListNode

	for p, cnt := this.head, -1; p != nil && p.next != nil; {
		if cnt+1 == index {
			prev = p
			break
		} else {
			p = p.next
		}
		cnt++
	}

	prev.next = prev.next.next
	if this.length == index {
		this.tail = prev
	}

	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestDesignLinkedList(t *testing.T) {

}
