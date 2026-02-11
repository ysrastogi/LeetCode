func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

func calc(root *TreeNode, sum *int) int {
    if root == nil {
        return 0
    }

    left := max(0, calc(root.Left, sum))
    right := max(0, calc(root.Right, sum))

    *sum = max(*sum, left+right+root.Val)

    return root.Val + max(left, right)
}

func maxPathSum(root *TreeNode) int {
    sum := root.Val
    calc(root, &sum)
    return sum
}
