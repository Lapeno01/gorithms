package bst

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func (rt *TreeNode) Insert(key int) {
	if key < rt.Key {
		if rt.Left == nil {
			rt.Left = &TreeNode{Key: key}
		} else {
			rt.Left.Insert(key)
		}
	} else if key > rt.Key {
		if rt.Right == nil {
			rt.Right = &TreeNode{Key: key}
		} else {
			rt.Right.Insert(key)
		}
	}
}

func (rt *TreeNode) Search(key int) bool {
	if rt == nil {
		return false
	}
	if key < rt.Key {
		return rt.Left.Search(key)
	} else if key > rt.Key {
		return rt.Right.Search(key)
	}
	return true
}

func (rt *TreeNode) Delete(key int) *TreeNode {
	if rt == nil {
		return nil
	}
	if key < rt.Key {
		rt.Left = rt.Left.Delete(key)
	} else if key > rt.Key {
		rt.Right = rt.Right.Delete(key)
	} else {
		if rt.Left == nil {
			return rt.Right
		} else if rt.Right == nil {
			return rt.Left
		}
		minRight := rt.Right.Min()
		rt.Key = minRight.Key
		rt.Right = rt.Right.Delete(rt.Key)
	}
	return rt
}

func (rt *TreeNode) Min() *TreeNode {
	current := rt
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}
