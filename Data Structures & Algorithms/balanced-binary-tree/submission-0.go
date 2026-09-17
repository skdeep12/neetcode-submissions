/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	_, result := recurse(root)
	return result
}

func recurse(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lh, lr := recurse(root.Left)
	rh,rr := recurse(root.Right)
	if lr && rr && abs(lh-rh) < 2 {
		return max(lh,rh)+1, true
	}
	return max(lh,rh)+1, false
}

func abs(r int) int{
	if r < 0{
		return -r
	}
	return r
}


