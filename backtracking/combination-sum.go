func backTrack(
    i int,
    curr []int,
    result *[][]int,
    candidates []int,
    sum int,
    target int,
) {
    if sum == target {
        temp := make([]int, len(curr))
        copy(temp, curr)
        *result = append(*result, temp)
        return
    }

    if i >= len(candidates) || sum > target {
        return
    }

    curr = append(curr, candidates[i])
    backTrack(i, curr, result, candidates, sum+candidates[i], target)
    curr = curr[:len(curr)-1]
    backTrack(i+1, curr, result, candidates, sum, target)
}

func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    backTrack(0, []int{}, &result, candidates, 0, target)
    return result
}
