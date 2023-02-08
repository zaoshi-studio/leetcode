package leetcode

import (
	"sort"
	"testing"
	"time"
)

/**
LeetCode company workers use key-cards to unlock office doors. Each time a
worker uses their key-card, the security system saves the worker's name and the time
when it was used. The system emits an alert if any worker uses the key-card
three or more times in a one-hour period.

 You are given a list of strings keyName and keyTime where [keyName[i], keyTime[
i]] corresponds to a person's name and the time when their key-card was used in
a single day.

 Access times are given in the 24-hour time format "HH:MM", such as "23:51" and
"09:49".

 Return a list of unique worker names who received an alert for frequent
keycard use. Sort the names in ascending order alphabetically.

 Notice that "10:00" - "11:00" is considered to be within a one-hour period,
while "22:51" - "23:52" is not considered to be within a one-hour period.


 Example 1:


Input: keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"],
keyTime = ["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]
Output: ["daniel"]
Explanation: "daniel" used the keycard 3 times in a one-hour period ("10:00","10
:40", "11:00").


 Example 2:


Input: keyName = ["alice","alice","alice","bob","bob","bob","bob"], keyTime = [
"12:01","12:00","18:00","21:00","21:20","21:30","23:00"]
Output: ["bob"]
Explanation: "bob" used the keycard 3 times in a one-hour period ("21:00","21:20
", "21:30").



 Constraints:


 1 <= keyName.length, keyTime.length <= 10⁵
 keyName.length == keyTime.length
 keyTime[i] is in the format "HH:MM".
 [keyName[i], keyTime[i]] is unique.
 1 <= keyName[i].length <= 10
 keyName[i] contains only lowercase English letters.


 Related Topics 数组 哈希表 字符串 排序 👍 68 👎 0

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
func alertNames(keyName []string, keyTime []string) []string {

	// 一小时时间内：指的是当天内任意一小时时间内，也即任意时间点之差小于等于60分钟
	var ans []string

	record := make(map[string][]int)

	for k, v := range keyTime {
		t, _ := time.Parse("15:04", v)
		record[keyName[k]] = append(record[keyName[k]], t.Hour()*60+t.Minute())
	}

	for name, times := range record {
		// 对一个员工的所有时间进行排序

		sort.Ints(times)

		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				ans = append(ans, name)
				break
			}
		}

	}

	sort.Strings(ans)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAlertUsingSameKeyCardThreeOrMoreTimesInAOneHourPeriod(t *testing.T) {

}
