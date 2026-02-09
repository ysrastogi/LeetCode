func lowerBound(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n-1
    ans := n
    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] >= target {
            ans = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return ans
}

func upperBound(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n-1
    ans := n
    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] > target {
            ans = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return ans
}

func searchRange(nums []int, target int) []int {
    n := len(nums)
    ans := []int{-1, -1}
    if n == 0 {
        return ans
    }

    lower := lowerBound(nums, target)
    if lower == n || nums[lower] != target {
        return ans
    }

    upper := upperBound(nums, target)
    ans[0] = lower
    ans[1] = upper - 1
    return ans
}
