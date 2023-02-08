package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/**
Given a list of folders folder, return the folders after removing all sub-
folders in those folders. You may return the answer in any order.

 If a folder[i] is located within another folder[j], it is called a sub-folder
of it.

 The format of a path is one or more concatenated strings of the form: '/'
followed by one or more lowercase English letters.


 For example, "/leetcode" and "/leetcode/problems" are valid paths while an
empty string and "/" are not.



 Example 1:


Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
Output: ["/a","/c/d","/c/f"]
Explanation: Folders "/a/b" is a subfolder of "/a" and "/c/d/e" is inside of
folder "/c/d" in our filesystem.


 Example 2:


Input: folder = ["/a","/a/b/c","/a/b/d"]
Output: ["/a"]
Explanation: Folders "/a/b/c" and "/a/b/d" will be removed because they are
subfolders of "/a".


 Example 3:


Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
Output: ["/a/b/c","/a/b/ca","/a/b/d"]



 Constraints:


 1 <= folder.length <= 4 * 10⁴
 2 <= folder[i].length <= 100
 folder[i] contains only lowercase letters and '/'.
 folder[i] always starts with the character '/'.
 Each folder name is unique.


 Related Topics 字典树 数组 字符串 👍 105 👎 0

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
func removeSubfolders1(folder []string) []string {

	sort.Strings(folder)

	var ans []string

	treeMap := make(map[string]struct{})

	for _, v := range folder {
		fnames := strings.Split(v, "/")

		target := "/"

		flag := false

		for _, vv := range fnames[1:] {
			target += vv
			if _, ok := treeMap[target]; ok {
				flag = true
				break
			}
			target += "/"
		}

		if !flag {
			treeMap[v] = struct{}{}
		}
	}

	for k := range treeMap {
		ans = append(ans, k)
	}

	return ans
}

// 排序 + 双指针
func removeSubfolders2(folder []string) []string {

	sort.Strings(folder)

	var ans []string

	// i,j 不相同表示找到了新的根文件夹
	// 仅当 i,j 相同时，表示从新的根文件夹开始比较，此时将该文件夹名称加入结果集
	// folder[i] + "/ 是为了解决 /a/b/c 和 /a/b/ca 这种问题
	for i, j := 0, 0; j < len(folder); j++ {
		if strings.HasPrefix(folder[j], folder[i]+"/") && i != j {
			continue
		}
		i = j
		ans = append(ans, folder[i])

	}

	return ans
}

// 字典树
// 类似 go 的路由框架中的请求路径匹配
type DictionaryTree struct {
	children map[string]*DictionaryTree
	isEnd    bool
}

func NewDictionaryTree() *DictionaryTree {
	return &DictionaryTree{
		children: map[string]*DictionaryTree{},
	}
}

func (this *DictionaryTree) Insert(str string) {

	node := this

	for _, v := range strings.Split(str, "/")[1:] {
		if _, ok := node.children[v]; !ok {
			node.children[v] = NewDictionaryTree()
		}
		node = node.children[v]
	}

	node.isEnd = true
}

func (this *DictionaryTree) Query(str string) bool {
	node := this

	for _, v := range strings.Split(str, "/")[1:] {
		if _, ok := node.children[v]; !ok {
			return false
		}

		node = node.children[v]
		if node.isEnd {
			return true
		}
	}

	return false
}

func removeSubfolders(folder []string) []string {

	sort.Strings(folder)

	var ans []string

	tree := NewDictionaryTree()

	for _, v := range folder {
		if !tree.Query(v) {
			tree.Insert(v)
			ans = append(ans, v)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveSubFoldersFromTheFilesystem(t *testing.T) {
	input := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}

	fmt.Println(removeSubfolders(input))
}
