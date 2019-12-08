package LinkedList

import (
	"fmt"
	"testing"
)

var l1 ItemLinkedList

func TestAdd(t *testing.T) {
	if !l1.IsEmpty() {
		t.Errorf("List should be empty")
	}

	l1.Add("first")
	if l1.IsEmpty() {
		t.Errorf("List should not be empty")
	}

	if size := l1.Size(); size != 1 {
		t.Errorf("wrong count. Expected 1 got %d", size)
	}

	l1.Add("Second")
	l1.Add("Third")

	if size := l1.Size(); size != 3 {
		t.Errorf("wrong count. Expected 3 got %d", size)
	}
}

func TestRemoveAt(t *testing.T) {
	_, err := l1.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpetcted error %s", err)
	}

	if size := l1.Size(); size != 2 {
		t.Errorf("wrong count. Expected 2 got %d", size)
	}
}

func TestAddAtIndex(t *testing.T) {
	err := l1.AddAtIndex(2, "second2")
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	if size := l1.Size(); size != 3 {
		t.Errorf("wrong count. Expected 3 got %d", size)
	}

	err = l1.AddAtIndex(0, "headTest")
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestIndexOf(t *testing.T) {
	if i := l1.IndexOf("headTest"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i := l1.IndexOf("first"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i := l1.IndexOf("second2"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i := l1.IndexOf("Third"); i != 3 {
		t.Errorf("expected position 3 but got %d", i)
	}
}

func TestHead(t *testing.T) {
	l1.Add("headTest")
	h := l1.Head()
	if "headTest" != fmt.Sprint(h.content) {
		t.Errorf("Expected `zero` but got %s", fmt.Sprint(h.content))
	}
}
