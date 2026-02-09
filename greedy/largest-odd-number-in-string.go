func largestOddNumber(num string) string {
    n := len(num)
        for i := n-1; i>=0; i-- {
            digit := num[i] - '0'
            if digit%2 == 1 {
                return num[:i+1]
            }
        }
        return ""
    
}