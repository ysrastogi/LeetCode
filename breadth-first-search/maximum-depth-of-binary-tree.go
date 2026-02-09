func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    height := 0
    q := []*TreeNode{root}

    for len(q) > 0 {
        currLen := len(q)
        for i := 0; i < currLen; i++ {
            node := q[0]
            q = q[1:]

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        height++
    }

    return height
}
