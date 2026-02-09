func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	ans := [][]int{}

	for len(q) > 0 {
		currLen := len(q)
		level := []int{}

		for i := 0; i < currLen; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ans = append(ans, level)
	}

	return ans
}
