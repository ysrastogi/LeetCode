func longestBalanced(nums []int) int {
    n := len(nums)
    maxLen := 0

    for i := 0; i < n; i++ {
        evenFreq := make(map[int]int)
        oddFreq := make(map[int]int)

        distinctEven := 0
        distinctOdd := 0

        for j := i; j < n; j++ {
            v := nums[j]

            if v%2 == 0 {
                if evenFreq[v] == 0 {
                    distinctEven++
                }
                evenFreq[v]++
            } else {
                if oddFreq[v] == 0 {
                    distinctOdd++
                }
                oddFreq[v]++
            }

            if distinctEven == distinctOdd {
                if j-i+1 > maxLen {
                    maxLen = j - i + 1
                }
            }
        }
    }

    return maxLen
}
