package binaryTree

import "fmt"

//定义节点数据结构
type BiTNode struct {
	Data   interface{}
	Lchild *BiTNode
	Rchild *BiTNode
}

type BiTree interface {
	Preorder(tree *BiTNode)
}

func (node *BiTNode) Preorder(tree *BiTNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data)
	node.Preorder(tree.Lchild)
	node.Preorder(tree.Rchild)
}

func (node BiTNode) Inorder(tree *BiTNode) {

}
