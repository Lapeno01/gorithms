package bst

func (rt *TreeNode) BFS() []int {
	if rt == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{rt}

	for len(queue) > 0 {
		// Dequeue the first element
		current := queue[0]
		queue = queue[1:]

		// Process the current node
		result = append(result, current.Key)

		// Enqueue left child if it exists
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		// Enqueue right child if it exists
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}
