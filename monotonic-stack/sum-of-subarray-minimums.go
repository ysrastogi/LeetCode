func findNSE(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = n
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func findPSE(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	pse := findPSE(arr)
	nse := findNSE(arr)
	res := 0
	mod := 1000000007

	for i := 0; i < n; i++ {
		left := i - pse[i]
		right := nse[i] - i
		res = (res + arr[i]*left*right) % mod
	}
	return res

    
}