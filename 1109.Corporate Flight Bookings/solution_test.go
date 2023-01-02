package leetcode

import (
	"testing"
)

/**
There are n flights that are labeled from 1 to n.

 You are given an array of flight bookings bookings, where bookings[i] = [
firsti, lasti, seatsi] represents a booking for flights firsti through lasti (
inclusive) with seatsi seats reserved for each flight in the range.

 Return an array answer of length n, where answer[i] is the total number of
seats reserved for flight i.


 Example 1:


Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
Output: [10,55,45,25,25]
Explanation:
Flight labels:        1   2   3   4   5
Booking 1 reserved:  10  10
Booking 2 reserved:      20  20
Booking 3 reserved:      25  25  25  25
Total seats:         10  55  45  25  25
Hence, answer = [10,55,45,25,25]


 Example 2:


Input: bookings = [[1,2,10],[2,2,15]], n = 2
Output: [10,25]
Explanation:
Flight labels:        1   2
Booking 1 reserved:  10  10
Booking 2 reserved:      15
Total seats:         10  25
Hence, answer = [10,25]




 Constraints:


 1 <= n <= 2 * 10⁴
 1 <= bookings.length <= 2 * 10⁴
 bookings[i].length == 3
 1 <= firsti <= lasti <= n
 1 <= seatsi <= 10⁴


 Related Topics 数组 前缀和 👍 430 👎 0

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
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	ans := make([]int, n+1)

	// 在差分的过程中，只用保证区间头尾的变化
	// 在对区间内的值变化时，差分数组的相对大小不会改变
	for _, arr := range bookings {
		diff[arr[0]] += arr[2]
		if arr[1]+1 < len(diff) {
			diff[arr[1]+1] -= arr[2]
		}
	}

	for i := 1; i < len(diff); i++ {
		ans[i] = ans[i-1] + diff[i]
	}
	return ans[1:]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCorporateFlightBookings(t *testing.T) {

}
