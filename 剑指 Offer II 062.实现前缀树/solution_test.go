package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 设计 字典树 哈希表 字符串 👍 44 👎 0

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
type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: map[byte]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	node := this

	for i := 0; i < len(word); i++ {
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[word[i]]
	}

	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	node := this

	for i := 0; i < len(word); i++ {
		if _, ok := node.children[word[i]]; !ok {
			return false
		}
		node = node.children[word[i]]
	}

	if node.isEnd {
		return true
	}

	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this

	for i := 0; i < len(prefix); i++ {
		if _, ok := node.children[prefix[i]]; !ok {
			return false
		}
		node = node.children[prefix[i]]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestQC3q1f(t *testing.T) {

}
