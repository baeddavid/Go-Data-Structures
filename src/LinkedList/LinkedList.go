package LinkedList

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	content Item
	next    *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (l1 *ItemLinkedList) Add(t Item) {
	l1.lock.Lock()
	node := Node{t, nil}
	if l1.head == nil {
		l1.head = &node
	} else {
		current := l1.head
		for {
			if current.next == nil {
				break
			}
			current = current.next
		}
		current.next = &node
	}
	l1.size++
	l1.lock.Unlock()
}

func (l1 *ItemLinkedList) AddAtIndex(i int, t Item) error {
	l1.lock.Lock()
	defer l1.lock.Unlock()

	if i < 0 || i > l1.size {
		return fmt.Errorf("Index out of bounds")
	}
	newNode := Node{t, nil}
	if i == 0 {
		newNode.next = l1.head
		l1.head = &newNode
		return nil
	}
	node := l1.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
	l1.size++
	return nil
}

func (li *ItemLinkedList) RemoveAt(i int) (*Item, error) {
	li.lock.Lock()
	defer li.lock.Unlock()
	if i < 0 || i > li.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	current := li.head
	j := 0
	for j < i-1 {
		j++
		current = current.next
	}
	deleted := current.next
	current.next = deleted.next
	li.size--
	return &current.content, nil
}

func (li *ItemLinkedList) IndexOf(t Item) int {
	li.lock.RLock()
	defer li.lock.RUnlock()
	current := li.head
	j := 0
	for {
		if current.content == t {
			return j
		}
		if current.next == nil {
			return -1
		}
		current = current.next
		j++
	}
}

func (li *ItemLinkedList) IsEmpty() bool {
	li.lock.RLock()
	defer li.lock.RUnlock()
	if li.head == nil {
		return true
	}
	return false
}

func (li *ItemLinkedList) Size() int {
	li.lock.RLock()
	defer li.lock.RUnlock()
	size := 1
	current := li.head
	for {
		if current == nil || current.next == nil {
			break
		}
		current = current.next
		size++
	}
	return size
}

func (li *ItemLinkedList) String() {
	li.lock.RLock()
	defer li.lock.RUnlock()
	current := li.head
	j := 0
	for {
		if current == nil {
			break
		}
		j++
		fmt.Print(current.content)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func (li *ItemLinkedList) Head() *Node {
	li.lock.RLock()
	defer li.lock.RUnlock()
	return li.head
}
