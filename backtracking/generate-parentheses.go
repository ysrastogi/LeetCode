func backTrack(curr string, open int, close int, n int, result *[]string) {
    if len(curr) == 2*n {
        *result = append(*result, curr)
        return
    }

    if open < n {
        backTrack(curr+"(", open+1, close, n, result)
    }

    if close < open {
        backTrack(curr+")", open, close+1, n, result)
    }
}

func generateParenthesis(n int) []string {
    result := []string{}
    backTrack("", 0, 0, n, &result)
    return result
}
