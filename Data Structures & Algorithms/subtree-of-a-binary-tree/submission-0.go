/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
    return inorder(root, subRoot)
}


func inorder(root *TreeNode, subRoot *TreeNode) bool{
	if root == nil {
		return false
	}
	if root.Val == subRoot.Val {
		if isSame(root,subRoot){
			return true
		}
	} 
	if inorder(root.Left, subRoot){
		return true
	}	
	return inorder(root.Right, subRoot)
}

func isSame(root *TreeNode, subRoot *TreeNode) bool{
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right){
		return true
	}
	return false
}