package Stack

import (
	"testing"
)

var s ItemStack

func InitStack() *ItemStack {
	s = ItemStack{}
	s.Stack(2)
	return &s
}

func TestPush(t *testing.T) {
	s := InitStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if size := s.Size(); size != 3 {
		t.Errorf("Wrong count expected 3 got %d", size)
	}
}

func TestPop(t *testing.T) {
	s.Pop()
	if size := s.Size(); size != 2 {
		t.Errorf("Wrong count expected 2 got %d", size)
	}
	s.Pop()
	s.Pop()
	if size := s.Size(); size != 0 {
		t.Errorf("Wrong count expected 0 got %d", size)
	}
}
