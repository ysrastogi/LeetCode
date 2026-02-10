func max(x int, y int)int{
    if x>y {
        return x
    }
    return y
}
func calc(root *TreeNode, diameter *int) int{
    if root == nil {
        return 0
    }
    lh := calc(root.Left, diameter)
    rh := calc(root.Right, diameter)
    *diameter = max(*diameter, lh+rh)
    return 1+max(lh, rh)
}
func diameterOfBinaryTree(root *TreeNode) int {
    diameter :=0
    calc(root, &diameter)
    return diameter
}