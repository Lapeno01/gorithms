package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() {
		t.Errorf("NewStack() should create an empty stack")
	}
	if s.Size() != 0 {
		t.Errorf("NewStack() should create a stack with size 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Push() failed, expected size 3, got %d", s.Size())
	}

	if s.IsEmpty() {
		t.Errorf("Push() failed, stack should not be empty")
	}
}

func TestPop(t *testing.T) {
	s := NewStack()

	// Test pop from empty stack
	_, ok := s.Pop()
	if ok {
		t.Errorf("Pop() from empty stack should return false")
	}

	// Test pop from non-empty stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, ok := s.Pop()
	if !ok || item != 3 {
		t.Errorf("Pop() failed, expected 3 and true, got %d and %v", item, ok)
	}

	if s.Size() != 2 {
		t.Errorf("Pop() failed, expected size 2 after pop, got %d", s.Size())
	}

	// Pop remaining items
	item, ok = s.Pop()
	if !ok || item != 2 {
		t.Errorf("Pop() failed, expected 2 and true, got %d and %v", item, ok)
	}

	item, ok = s.Pop()
	if !ok || item != 1 {
		t.Errorf("Pop() failed, expected 1 and true, got %d and %v", item, ok)
	}

	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after popping all elements")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack()

	// Test peek from empty stack
	_, ok := s.Peek()
	if ok {
		t.Errorf("Peek() from empty stack should return false")
	}

	// Test peek from non-empty stack
	s.Push(1)
	s.Push(2)

	item, ok := s.Peek()
	if !ok || item != 2 {
		t.Errorf("Peek() failed, expected 2 and true, got %d and %v", item, ok)
	}

	// Size should remain the same after peek
	if s.Size() != 2 {
		t.Errorf("Peek() should not change stack size, expected 2, got %d", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Errorf("IsEmpty() failed, new stack should be empty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() failed, stack with elements should not be empty")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() failed, stack should be empty after popping all elements")
	}
}

func TestSize(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf("Size() failed, empty stack should have size 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Size() failed, expected 3, got %d", s.Size())
	}

	s.Pop()
	if s.Size() != 2 {
		t.Errorf("Size() failed, expected 2 after pop, got %d", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if !s.IsEmpty() {
		t.Errorf("Clear() failed, stack should be empty")
	}

	if s.Size() != 0 {
		t.Errorf("Clear() failed, stack size should be 0, got %d", s.Size())
	}
}

func TestToSlice(t *testing.T) {
	s := NewStack()

	// Test empty stack
	slice := s.ToSlice()
	if len(slice) != 0 {
		t.Errorf("ToSlice() failed, empty stack should return empty slice, got %v", slice)
	}

	// Test non-empty stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	slice = s.ToSlice()
	expected := []int{3, 2, 1} // top to bottom

	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("ToSlice() failed, expected %v, got %v", expected, slice)
	}
}

func TestStackOperations(t *testing.T) {
	s := NewStack()

	// Test LIFO behavior
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Should pop in reverse order (LIFO)
	item1, _ := s.Pop()
	item2, _ := s.Pop()
	item3, _ := s.Pop()

	if item1 != 30 || item2 != 20 || item3 != 10 {
		t.Errorf("Stack LIFO behavior failed, expected 30,20,10 got %d,%d,%d", item1, item2, item3)
	}
}

func TestStackPeekAfterOperations(t *testing.T) {
	s := NewStack()
	s.Push(100)
	s.Push(200)

	// Peek should show top element
	top, ok := s.Peek()
	if !ok || top != 200 {
		t.Errorf("Peek() failed after push operations, expected 200, got %d", top)
	}

	// Pop and peek again
	s.Pop()
	top, ok = s.Peek()
	if !ok || top != 100 {
		t.Errorf("Peek() failed after pop operation, expected 100, got %d", top)
	}
}
