func search(nums []int, target int) int {
    l, h := 0, len(nums)-1

    for l <= h {
        mid := l + (h-l)/2

        if nums[mid] == target {
            return mid
        }

        // Left half is sorted
        if nums[l] <= nums[mid] {
            if nums[l] <= target && target < nums[mid] {
                h = mid - 1
            } else {
                l = mid + 1
            }
        } else { // Right half is sorted
            if nums[mid] < target && target <= nums[h] {
                l = mid + 1
            } else {
                h = mid - 1
            }
        }
    }

    return -1
}
