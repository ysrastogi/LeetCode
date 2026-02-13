/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import(
     "container/list"
     "sort"
     )

type item struct {
	node *TreeNode
	x    int
	y    int
}

func verticalTraversal(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}

	nodes := map[int]map[int][]int{}

	q := list.New()
	q.PushBack(item{root, 0, 0})

	for q.Len() > 0 {
		front := q.Front()
		q.Remove(front)

		cur := front.Value.(item)
		n, x, y := cur.node, cur.x, cur.y

		if _, ok := nodes[x]; !ok {
			nodes[x] = map[int][]int{}
		}
		nodes[x][y] = append(nodes[x][y], n.Val)

		if n.Left != nil {
			q.PushBack(item{n.Left, x - 1, y + 1})
		}
		if n.Right != nil {
			q.PushBack(item{n.Right, x + 1, y + 1})
		}
	}

	xKeys := make([]int, 0, len(nodes))
	for k := range nodes {
		xKeys = append(xKeys, k)
	}
	sort.Ints(xKeys)

	ans := [][]int{}

	for _, x := range xKeys {
		col := []int{}

		yMap := nodes[x]
		yKeys := make([]int, 0, len(yMap))
		for k := range yMap {
			yKeys = append(yKeys, k)
		}
		sort.Ints(yKeys)

		for _, y := range yKeys {
			vals := yMap[y]
			sort.Ints(vals)
			col = append(col, vals...)
		}

		ans = append(ans, col)
	}

	return ans
    
}