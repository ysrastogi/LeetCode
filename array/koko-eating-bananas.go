func minEatingSpeed(piles []int, h int) int {
    lo := 1
    hi := 0

    // Find max pile
    for _, p := range piles {
        if p > hi {
            hi = p
        }
    }

    ans := hi

    for lo <= hi {
        mid := lo + (hi-lo)/2

        if canEat(piles, mid, h) {
            ans = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }

    return ans
}

func canEat(piles []int, k int, h int) bool {
    hours := 0

    for _, p := range piles {
        hours += (p + k - 1) / k
        if hours > h {
            return false
        }
    }

    return true
}
