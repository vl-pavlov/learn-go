package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println("\nSmall Message")
	fmt.Println("")

	message := []byte("Data structure is a data organization, and storage format that is usually chosen for efficient access to data.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n\nLarge Message")
	fmt.Println("")

	message = []byte("In computer science, a red–black tree is a specialised binary search tree data structure noted for fast storage and retrieval of ordered information, and a guarantee that operations will complete within a known time. Compared to other self-balancing binary search trees, the nodes in a red-black tree hold an extra bit called \"color\" representing \"red\" and \"black\" which is used when re-organising the tree to ensure that it is always approximately balanced.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}
