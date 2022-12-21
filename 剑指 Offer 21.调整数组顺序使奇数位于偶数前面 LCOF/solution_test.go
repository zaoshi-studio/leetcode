package leetcode

import (
	"sort"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 数组 双指针 排序 👍 270 👎 0

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
func exchange1(nums []int) []int {
	var (
		even []int
		odd  []int
	)

	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return append(odd, even...)
}

func exchange2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i]%2 != 0 {
			return true
		}
		return false
	})
	return nums
}

func exchange3(nums []int) []int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	ans := make([]int, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			ans[right] = v
			right--
		} else {
			ans[left] = v
			left++
		}
	}
	return ans
}

func exchange(nums []int) []int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right {

		for ; left < right && nums[left]%2 != 0; left++ {
		}

		for ; left < right && nums[right]%2 == 0; right-- {
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDiaoZhengShuZuShunXuShiQiShuWeiYuOuShuQianMianLcof(t *testing.T) {

}
