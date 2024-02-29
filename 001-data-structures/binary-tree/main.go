package main

import (
	"fmt"
	"io"
	"os"
)

// BinaryNode represents a node in a binary tree with a left and right child node and an integer data value.
type BinaryNode struct {
	left  *BinaryNode // Pointer to the left child node
	right *BinaryNode // Pointer to the right child node
	data  int         // Integer data value stored in the node
}

// BinaryTree represents a binary tree with a root node.
type BinaryTree struct {
	root *BinaryNode // Pointer to the root node of the binary tree
}

// insert inserts a new integer value into the binary tree.
func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// insert inserts a new integer value into the binary tree, recursively traversing the tree to find the correct location.
func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

// print prints the binary tree in a human-readable format, with indentation and labels for each node.
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

// main is the entry point of the program, where the binary tree is constructed and printed.
func main() {
	tree := &BinaryTree{}

	tree.insert(50).
		insert(-20).
		insert(-50).
		insert(-60).
		insert(30).
		insert(60).
		insert(55).
		insert(85).
		insert(5).
		insert(-10).
		insert(100)

	print(os.Stdout, tree.root, 0, 'M')
}
