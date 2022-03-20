package main

import "struct/binaryTree"

func main() {
	//root := binaryTree.BiTNode{}
	//left := binaryTree.BiTNode{}
	//right := binaryTree.BiTNode{}
	//root.Data = 1
	//left.Data = 2
	//right.Data = 3
	//root.Lchild = &left
	//root.Rchild = &right
	//root.Preorder(&root)
	var a, b, c binaryTree.HuTNode
	a.Weight = 1
	a.Value = "a"
	b.Weight = 2
	b.Value = "b"
	c.Value = "c"
	c.Weight = 3
	binaryTree.HuffmanEncode([]binaryTree.HuTNode{a, b, c})

}
