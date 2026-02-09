func traversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)

	traversal(root.Left, result)
	traversal(root.Right, result)
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	traversal(root, &result)
	return result
}
