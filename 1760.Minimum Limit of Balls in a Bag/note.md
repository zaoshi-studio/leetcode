# 1760.袋子里最少数目的球【如何理解官方题解：个人理解】

```go
func minimumSize(nums []int, maxOperations int) int {
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}

	return sort.Search(max, func(target int) bool {
		if target == 0 {
			return false
		}
		ops := 0
		for _, x := range nums {
			ops += (x - 1) / target
		}
		return ops <= maxOperations
	})
}

```

这里对官方题解做了点小小的修改，`sort.Search()` 中回调函数的形参由 `y` 改为 `target`

## 理解 sort.Search

```go
func Search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1 
		} else {
			j = h 
		}
	}
	return i
}
```

通过查看源码，可以看到 `sort.Search` 就是一个 **二分查找** 的模板，但需要外部传入回调函数确定 **搜索范围的收缩方向**
> 回调函数中的 `i`,`j` 是实际含有大小的值，而不是通常二分查找中的下标

回调函数中的参数 `h` 就是边界 `i` 和 `j` 的中间值

- 当回调函数返回 `false` 则向右收缩 `[h+1, j]`
- 当回调函数返回 `true` 则向左收缩 `[i,h]`

---

## 理解官方题解中的思路

对于给定的数组 `nums[]` 加入没有 `maxOperations` 的限制，则我们一定可以将所有的数拆分成若干个 `1`

即最后的袋子一定会变成 `nums[]int{1,1,1, ... , 1}`

因此最终的解一定落在 `[1: max(nums)]` 中

**`Search` 函数中的二分实际上是在 `[1: max(nums)]` 区间内进行二分查找**


在确定了查找范围后，如何判定查找到的值是满足条件的，这就需要定义回调函数

回调函数中的 `target` 就是我们每轮循环中假定的解

这个解一定不为 `0`
- 我们不能将每个袋子中的球分成 0 个
- 涉及到后续计算，分母也不能为 0

---

### 理解 ops

`ops` 的作用是计算 **假设要将当前数组所有的值划分成各部分 $\le$ `target` 的值所需要进行的操作次数**


---

### 理解操作次数计算公式

假设有一个数 `num` 要将其划分成各部分 $\le$ `div` ，需要多少次操作

最简单的做法直接相除 `num/div`

但如果 `num` 与 `div` 可以直接整除，则我们会多算一次，所需要的操作数被 `+1`
此处的 `+1` 是包含 **除到 0** 的操作步骤
> 我们仅需要各部分 $\le$ `div` 也就是当剩余部分 = `div` 时，不需要在进行分割

因此在除之前进行 `-1` 操作，冗余一下，即使 `-1` 后恰好能够整除 `(num - 1)/div = 0` 则表明多余的 `1` 仍要进行一次分割


### 理解回调函数

每当通过二分查找找到一个 **分割值 `target`** 我们就进行一次计算

计算数组中所有的数，分割到满足条件所需要的操作总数

如果满足，则缩小划分范围(减少 `target`)；如果不满足，则要扩大划分范围(增加 `target`)
> `target` 越大进行分割的次数越少