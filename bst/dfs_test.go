package bst

import (
	"reflect"
	"testing"
)

func TestDFSInOrder(t *testing.T) {
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

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	result := root.DFSInOrder()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DFSInOrder() = %v; want %v", result, expected)
	}
}
