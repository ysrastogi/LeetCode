/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traversal(root *TreeNode, result *[]int){
    if root==nil{
        return
    }
    traversal(root.Left, result)
    *result = append(*result, root.Val)
    traversal(root.Right, result)
}
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    traversal(root, &result)
    return result
    
}