/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := make(map[int]*Node)
	if node == nil {
		return nil
	}
	return dfs(node, visited)

}

func dfs(node *Node, visited map[int]*Node) *Node{
	if val, ok := visited[node.Val]; ok {
		return val
	} else {
		n := Node{node.Val, make([]*Node, len(node.Neighbors))}
		visited[node.Val] = &n
		for i:=0;i<len(node.Neighbors);i+=1{
			n.Neighbors[i] = dfs(node.Neighbors[i], visited)
		}
		return &n
	}
}
