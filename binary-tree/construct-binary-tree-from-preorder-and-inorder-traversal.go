var preIdx int

func build(
    inMap map[int]int,
    inStart int,
    inEnd int,
    preorder []int,
) *TreeNode {

    if inStart > inEnd {
        return nil
    }

    rootVal := preorder[preIdx]
    preIdx++

    root := &TreeNode{Val: rootVal}
    mid := inMap[rootVal]
    root.Left = build(inMap, inStart, mid-1, preorder)
    root.Right = build(inMap, mid+1, inEnd, preorder)

    return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    inMap := make(map[int]int)
    for i, v := range inorder {
        inMap[v] = i
    }

    preIdx = 0
    return build(inMap, 0, len(inorder)-1, preorder)
}
