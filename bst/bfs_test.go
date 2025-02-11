package bst

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	/*
	       4
	      / \
	     2   6
	    / \ / \
	   1  3 5  7
	*/
	root := &TreeNode{Key: 4}
	root.Insert(2)
	root.Insert(6)
	root.Insert(1)
	root.Insert(3)
	root.Insert(5)
	root.Insert(7)

	expected := []int{4, 2, 6, 1, 3, 5, 7}
	result := root.BFS()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS() = %v; want %v", result, expected)
	}
}
