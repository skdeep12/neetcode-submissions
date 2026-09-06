/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func goodNodes(root *TreeNode) int {
	ans = 0
	traverse(root, root.Val)
	return ans
}

func traverse(root *TreeNode, m int) {
	if root == nil {
		return
	}
	if m <= root.Val {
		ans +=1 
	}
	m = max(m,root.Val)
	traverse(root.Left,m)
	traverse(root.Right, m)
}