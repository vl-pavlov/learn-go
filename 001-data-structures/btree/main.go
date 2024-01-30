package main

import (
	"fmt"

	"github.com/google/btree"
)

const (
	degree = 10
)

func main() {
	tree := btree.NewOrderedG[btree.Int](degree)

	tree.ReplaceOrInsert(btree.Int(10))
	tree.ReplaceOrInsert(btree.Int(15))
	tree.ReplaceOrInsert(btree.Int(30))

	fmt.Println(tree.Get(btree.Int(10)))
	// Output: 10 true

	fmt.Println(tree.Get(btree.Int(1000)))
	// Output: 0 false

}
