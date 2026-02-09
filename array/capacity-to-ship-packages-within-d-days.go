func canShip(weights []int, days int, total_q int) bool {
    req_d := 1
    total_q_temp := total_q

    for i := 0; i < len(weights); i++ {
        if weights[i] > total_q {
            return false
        }

        if weights[i] <= total_q_temp {
            total_q_temp -= weights[i]
        } else {
            req_d++
            total_q_temp = total_q - weights[i]
        }
    }

    return req_d <= days
}

func shipWithinDays(weights []int, days int) int {
    min_q := 0
    max_q := 0

    for _, w := range weights {
        if w > min_q {
            min_q = w
        }
        max_q += w
    }

    ans := max_q

    for min_q <= max_q {
        mid := min_q + (max_q-min_q)/2

        if canShip(weights, days, mid) {
            ans = mid
            max_q = mid - 1
        } else {
            min_q = mid + 1
        }
    }

    return ans
}
