func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    slices.Sort(strs)

    first := strs[0]
    last := strs[len(strs)-1]

    minLen := min(len(first), len(last))
    i := 0

    for i < minLen && first[i] == last[i] {
        i++
    }

    return first[:i]
}
