```go
func threeSum(nums []int) [][]int {
	var res [][]int 
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			cur := nums[i] + nums[j] + nums[k]
			if 0 == cur {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
			if cur > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
```

仍是采用 **双指针法**，将 i 作为固定遍历的指针，实际的双指针为 **j** 和 **k**

当前代码仅能找出所有满足条件的解，但无法完成 **去重**

要实现 **去重** 只能从三个指针入手

## 对指针 i 去重

指针 i 去重的方式有两种

```go
if nums[i] == nums[i+1]{
	continue
}
```

**(一)** 使用该方法进行去重，当一个元素检测到后续有与之相同的元素后会不断跳过， **直到当前元素仅保留一个时** 才会进入后续的运算              

```go
if nums[i] == nums[i-1]{
	continue
}
```
**(二)** 使用该方法进行去重，当第一次碰到重复元素时正常进入运算，且会将重复的元素也纳入计算过程，这是符合结果的，例如 `[-1,-1,2]`


## 对指针 j 和 k 的去重

由于初始值的设置，`j = i + 1` 和 `k = len(nums) - 1` 三个指针 **在位置上** 必定不重复

当指针 `i` 在外轮循环固定之后，内部循环相当于双指针去找 `target = 0 - nums[i]` 的过程

对于一轮中固定的 `target` 每个解的 `{nums[j], nums[k]}` 必定不重复

对于在对指针 `k` 去重时要先 `--` 操作，`for k--; j < k && nums[k] == nums[k+1]; k-- {}`，指针 `k` 在每一轮循环中都被设置在末尾，直接用
`nums[k] == nums[k+1]` 判断会越界，`--` 操作后，如果相同则会继续循环，如果不相同则相当于完成了一次去重