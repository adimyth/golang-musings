package datastructures

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = bst.insert(node.Left, value)
	} else {
		node.Right = bst.insert(node.Right, value)
	}
	return node
}

func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = bst.insert(bst.Root, value)
}

func (bst *BinarySearchTree) search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return bst.search(node.Left, value)
	} else {
		return bst.search(node.Right, value)
	}
}

func (bst *BinarySearchTree) Search(value int) bool {
	return bst.search(bst.Root, value)
}

// InOrderTraversal: At first traverse left subtree then visit the root and then traverse the right subtree.
func (bst *BinarySearchTree) traversal(node *Node) {
	if node == nil {
		return
	}

	bst.traversal(node.Left)
	fmt.Printf("%d -> ", node.Value)
	bst.traversal(node.Right)
}

func (bst *BinarySearchTree) Traversal() {
	bst.traversal(bst.Root)
	fmt.Println()
}

// PreOrderTraversal: At first visit the root then traverse the left subtree and then traverse the right subtree.
func (bst *BinarySearchTree) preOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d -> ", node.Value)
	bst.preOrderTraversal(node.Left)
	bst.preOrderTraversal(node.Right)
}

func (bst *BinarySearchTree) PreOrderTraversal() {
	bst.preOrderTraversal(bst.Root)
	fmt.Println()
}

// PostOrderTraversal: At first traverse left subtree then traverse the right subtree and then visit the root.
func (bst *BinarySearchTree) postOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	bst.postOrderTraversal(node.Left)
	bst.postOrderTraversal(node.Right)
	fmt.Printf("%d -> ", node.Value)
}

func (bst *BinarySearchTree) PostOrderTraversal() {
	bst.postOrderTraversal(bst.Root)
	fmt.Println()
}

// InOrderSuccessor: The in-order successor of a node is the next node in the in-order traversal of the tree.
// The in-order successor of a node is the leftmost node in the right subtree of the node.
// 🔥 Replacing the node with its in-order successor will not change the in-order traversal of the tree since there are no nodes between the node and its in-order successor.
func (bst *BinarySearchTree) inOrderSuccessor(node *Node) *Node {
	current := node.Right

	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BinarySearchTree) delete(node *Node, value int) *Node {

	// If the tree is empty
	if node == nil {
		return node
	}

	// If the value to be deleted is less than the root value, then it lies in the left subtree
	if value < node.Value {
		node.Left = bst.delete(node.Left, value)
	}

	// If the value to be deleted is greater than the root value, then it lies in the right subtree
	if value > node.Value {
		node.Right = bst.delete(node.Right, value)
	}

	// If the value to be deleted is equal to the root value, then this is the node to be deleted
	if value == node.Value {

		// If the node has only one child or no child
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		// If the node has two children, then get the in-order successor (smallest in the right subtree)
		inOrderSuccessor := bst.inOrderSuccessor(node)

		// Copy the in-order successor value to the node
		node.Value = inOrderSuccessor.Value

		// Delete the in-order successor
		node.Right = bst.delete(node.Right, inOrderSuccessor.Value)
	}
	return node
}

func (bst *BinarySearchTree) Delete(value int) {
	bst.Root = bst.delete(bst.Root, value)
}
