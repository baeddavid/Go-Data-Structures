package Stack

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemStack struct {
	top       int
	nItems    int
	itemStack []Item
	lock      sync.RWMutex
}

func (s *ItemStack) Stack(INITIAL_SIZE int) *ItemStack {
	s.top = -1
	s.nItems = 0
	s.itemStack = make([]Item, INITIAL_SIZE, INITIAL_SIZE)
	return s
}

func (s *ItemStack) Push(item Item) {
	s.lock.Lock()
	s.top++
	if s.top == len(s.itemStack) {
		s.itemStack = append(s.itemStack, item)
		s.top++
	} else {
		s.itemStack[s.top] = item
	}
	s.nItems++
	s.lock.Unlock()
}

func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	if s.top == -1 {
		fmt.Println("ERROR STACK IS EMPTY")
	}
	temp := s.itemStack[s.top-1]
	s.top--
	s.nItems--
	s.lock.Unlock()
	return temp
}

func (s *ItemStack) Peek() Item {
	return s.itemStack[s.top-1]
}

func (s *ItemStack) Size() int {
	return s.nItems
}

func (s *ItemStack) Print() {
	for i := 0; i < s.top; i++ {
		fmt.Println(s.itemStack[i])
	}
}
