func reverse(nums []int) {
    l := 0
    r := len(nums) - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}

func nextPermutation(nums []int) {
    n := len(nums)
    ind := -1
    for i := n - 1; i > 0; i-- {
        if nums[i-1] < nums[i] {
            ind = i
            break
        }
    }
    if ind == -1 {
        reverse(nums)
        return
    }
    for i := n - 1; i >= ind; i-- {
        if nums[i] > nums[ind-1] {
            nums[i], nums[ind-1] = nums[ind-1], nums[i]
            break
        }
    }
    reverse(nums[ind:])
}