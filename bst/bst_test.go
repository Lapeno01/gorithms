package bst

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	root := &TreeNode{Key: 50}
	keys := []int{30, 70, 20, 40, 60, 80}

	for _, key := range keys {
		root.Insert(key)
	}

	tests := []struct {
		key      int
		expected bool
	}{
		{50, true},
		{30, true},
		{70, true},
		{20, true},
		{40, true},
		{60, true},
		{80, true},
		{25, false},
		{90, false},
	}

	for _, tt := range tests {
		result := root.Search(tt.key)
		if result != tt.expected {
			t.Errorf("Search(%d) = %v; want %v", tt.key, result, tt.expected)
		}
	}
}

func TestDelete(t *testing.T) {
	root := &TreeNode{Key: 50}
	keys := []int{30, 70, 20, 40, 60, 80}

	for _, key := range keys {
		root.Insert(key)
	}

	root = root.Delete(20) // Leaf node
	if root.Search(20) {
		t.Errorf("Delete(20) failed; 20 still found in BST")
	}

	root = root.Delete(30) // Node with one child
	if root.Search(30) {
		t.Errorf("Delete(30) failed; 30 still found in BST")
	}

	root = root.Delete(50) // Node with two children
	if root.Search(50) {
		t.Errorf("Delete(50) failed; 50 still found in BST")
	}
}
