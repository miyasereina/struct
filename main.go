package main

import "struct/binaryTree"

func main() {
	root := binaryTree.BiTNode{}
	left := binaryTree.BiTNode{}
	right := binaryTree.BiTNode{}
	root.Data = 1
	left.Data = 2
	right.Data = 3
	root.Lchild = &left
	root.Rchild = &right
	root.Preorder(&root)

}
