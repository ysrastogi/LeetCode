func previousSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	st := []int{}

	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[st[len(st)-1]] > nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = -1
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func nextSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	st := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] >= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = n
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func previousGreater(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	st := []int{}

	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = -1
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func nextGreater(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	st := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			res[i] = n
		} else {
			res[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func sumSubarrayMins(nums []int) int64 {
	n := len(nums)
	pse := previousSmaller(nums)
	nse := nextSmaller(nums)

	var total int64 = 0

	for i := 0; i < n; i++ {
		left := i - pse[i]
		right := nse[i] - i
		total += int64(nums[i]) * int64(left*right)
	}
	return total
}

func sumSubarrayMaxs(nums []int) int64 {
	n := len(nums)
	pge := previousGreater(nums)
	nge := nextGreater(nums)

	var total int64 = 0

	for i := 0; i < n; i++ {
		left := i - pge[i]
		right := nge[i] - i
		total += int64(nums[i]) * int64(left*right)
	}
	return total
}

func subArrayRanges(nums []int) int64 {
	return sumSubarrayMaxs(nums) - sumSubarrayMins(nums)
}
