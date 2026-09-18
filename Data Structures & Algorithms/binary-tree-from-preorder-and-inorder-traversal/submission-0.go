/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	val := preorder[0]
    node := TreeNode{val,nil,nil}
	preorder = preorder[1:]
	if len(preorder) == 0{
		return &node
	}
	for i:=0;i<len(inorder);i+=1{
		if inorder[i] == val {
			left := inorder[:i]
			right := inorder[i+1:]
			if len(left) > 0 {
				node.Left = buildTree(preorder[:i], left)
			}
			if len(right) > 0{
				node.Right = buildTree(preorder[i:], right)
			}
			break
		}
	}
	return &node

}
