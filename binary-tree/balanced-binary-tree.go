func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func dfsHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := dfsHeight(root.Left)
    if left == -1 {
        return -1
    }

    right := dfsHeight(root.Right)
    if right == -1 {
        return -1
    }

    if abs(left-right) > 1 {
        return -1
    }

    return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
    return dfsHeight(root) != -1
}
