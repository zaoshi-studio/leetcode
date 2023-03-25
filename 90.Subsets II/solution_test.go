package leetcode

import (
    "testing"
)

/**
Given an integer array nums that may contain duplicates, return all possible
subsets (the power set).

 The solution set must not contain duplicate subsets. Return the solution in
any order.


 Example 1:
 Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

 Example 2:
 Input: nums = [0]
Output: [[],[0]]


 Constraints:


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10


 Related Topics 位运算 数组 回溯 👍 1058 👎 0

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
func subsetsWithDup(nums []int) [][]int {
    var (
        ans [][]int
        n   = len(nums)
    )

    var backtrack func(idx int, buffer []int)
    backtrack = func(idx int, buffer []int) {
        tmp := make([]int, len(buffer))
        copy(tmp, buffer)
        ans = append(ans, tmp)

        for i := idx; i < n; i++ {
            if i > idx && nums[i] == nums[i-1] {
                continue
            }

            buffer = append(buffer, nums[i])
            backtrack(i+1, buffer)
            buffer = buffer[:len(buffer)-1]
        }
    }

    backtrack(0, []int{})

    return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubsetsIi(t *testing.T) {

}
