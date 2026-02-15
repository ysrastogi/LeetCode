/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leftHeight(node *TreeNode) int {
    h := 0
    for node != nil {
        h++
        node = node.Left
    }
    return h
}

func rightHeight(node *TreeNode) int {
    h := 0
    for node != nil {
        h++
        node = node.Right
    }
    return h
}
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh := leftHeight(root)
    rh := rightHeight(root)

    if lh == rh {
        return (1 << lh) - 1
    }

    return 1 + countNodes(root.Left) + countNodes(root.Right)
}