package linkedlist_double

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	var head *ListNode
	head = Append(head, 1)
	head = Append(head, 2)
	if head.Val != 1 || head.Next.Val != 2 || head.Next.Prev != head {
		t.Errorf("Append failed, expected 1->2 with proper prev links")
	}

	head = Prepend(head, 0)
	if head.Val != 0 || head.Next.Val != 1 || head.Next.Prev != head {
		t.Errorf("Prepend failed, expected 0->1->2 with proper prev links")
	}

	head = Delete(head, 1)
	current := head
	var values []int
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	expected := []int{0, 2}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Delete failed, expected %v, got %v", expected, values)
	}

	head = Append(head, 3)
	current = head
	values = []int{}
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	expected = []int{0, 2, 3}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Display sequence check failed, expected %v, got %v", expected, values)
	}

	current = head
	for current != nil && current.Next != nil {
		current = current.Next
	}

	for current != nil {
		if current.Next != nil && current.Next.Prev != current {
			t.Errorf("Prev link check failed at value %d", current.Val)
		}
		current = current.Prev
	}
}
