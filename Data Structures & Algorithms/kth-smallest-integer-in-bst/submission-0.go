/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
		return 0
	}
	ans, _ := traverse(root, 0, k)
	return ans
}

func traverse(root *TreeNode, k, K int) (int, int){
	if root == nil {
		return -1, k
	}
	l, k := traverse(root.Left, k, K)
	if l != -1 {
		return l, k
	}
	k+=1
	if k == K {
		return root.Val, k
	}
	return traverse(root.Right, k, K)
}
