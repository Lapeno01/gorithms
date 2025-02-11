package bst

func (rt *TreeNode) DFSInOrder() []int {
	var result []int
	dfs(rt, &result)
	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, result)
	*result = append(*result, node.Key)
	dfs(node.Right, result)
}
