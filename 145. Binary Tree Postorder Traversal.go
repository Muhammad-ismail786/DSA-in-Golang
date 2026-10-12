/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	postOrder(root, &res)
	return res

}
func postOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, res)
	postOrder(node.Right, res)
	*res = append(*res, node.Val)
}
