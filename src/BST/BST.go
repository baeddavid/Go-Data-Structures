package BST

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	key     int
	content Item
	left    *Node
	right   *Node
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *ItemBinarySearchTree) Add(key int, content Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{key, content, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		helper(bst.root, n)
	}
}

func helper(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			helper(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			helper(node.right, newNode)
		}
	}
}

func (bst *ItemBinarySearchTree) InOrderTrav(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTrav(bst.root, f)
}

func inOrderTrav(n *Node, f func(Item)) {
	if n != nil {
		inOrderTrav(n.left, f)
		f(n.content)
		inOrderTrav(n.right, f)
	}
}

func (bst *ItemBinarySearchTree) Getgi() *Item {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.left == nil {
			return &n.content
		}
		n = n.left
	}
}

func (bst *ItemBinarySearchTree) GetMax() *Item {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.right == nil {
			return &n.content
		}
		n = n.right
	}
}

func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return searchHelp(bst.root, key)
}

func searchHelp(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return searchHelp(n.left, key)
	}
	if key > n.key {
		return searchHelp(n.right, key)
	}
	return true
}

func (bst *ItemBinarySearchTree) Delete(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	deleteHelp(bst.root, key)
}

func deleteHelp(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = deleteHelp(node.left, key)
		return node
	}
	if key > node.key {
		node.right = deleteHelp(node.right, key)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftNode := node.right
	for {
		if leftNode != nil && leftNode.left != nil {
			leftNode = leftNode.left
		} else {
			break
		}
	}
	node.key, node.content = leftNode.key, leftNode.content
	node.right = deleteHelp(node.right, node.key)
	return node
}

func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}
