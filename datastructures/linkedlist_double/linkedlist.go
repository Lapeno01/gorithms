package linkedlist_double

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// Append adds a new node with the given value at the end of the list
func Append(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}
	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	newNode.Prev = current
	return head
}

// Prepend adds a new node with the given value at the beginning of the list
func Prepend(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value, Next: head}
	if head != nil {
		head.Prev = newNode
	}

	return newNode
}

// Delete removes the first occurrence of a node with the given value
func Delete(head *ListNode, value int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == value {
		newHead := head.Next
		if newHead != nil {
			newHead.Prev = nil
		}
		return newHead
	}

	current := head
	for current.Next != nil {
		if current.Next.Val == value {
			current.Next = current.Next.Next
			if current.Next != nil {
				current.Next.Prev = current
			}
			return head
		}
		current = current.Next
	}

	return head
}

// Display prints all values in the linked list
func Display(head *ListNode) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}

	fmt.Println("nil")
}
