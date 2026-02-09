func nextGreaterElements(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    for i:=0; i<n; i++{
        res[i]=-1
    }
    stack := []int{}

    for i:=0; i<2*n; i++{
        currIdx := i%n

        for len(stack)>0 && nums[currIdx] > nums[stack[len(stack)-1]]{
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[top] = nums[currIdx]
        }
        if i < n {
			stack = append(stack, currIdx)
		}
    }
    return res
}