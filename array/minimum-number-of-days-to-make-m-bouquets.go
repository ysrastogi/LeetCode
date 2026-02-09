func canMake(bloomDay []int, m int, k int, d int) bool {
    count := 0      // consecutive bloomed flowers
    bouquets := 0

    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] <= d {
            count++
            if count == k {
                bouquets++
                count = 0
            }
        } else {
            count = 0
        }
    }

    return bouquets >= m
}

func minDays(bloomDay []int, m int, k int) int {
    n:=len(bloomDay)
    if k*m>n{
        return -1
    }
    max_d:=0
    for i:=0; i<n; i++{
        if bloomDay[i]>max_d{
            max_d=bloomDay[i]
        }
    }
    lo:=0
    hi:=max_d
    ans:=0
    for lo<=hi{
        mid := lo+(hi-lo)/2
        if canMake(bloomDay, m, k, mid){
            ans = mid
            hi= mid-1
        }else{
            lo = mid+1
        }
    }
    return ans
}