package binaryTree

import (
	"fmt"
	"sort"
)

type HuTNode struct {
	Weight int
	Value  interface{}
	Parent *HuTNode
	Lchild *HuTNode
	Rchild *HuTNode
}

var HTCodes map[interface{}]string

func HuffmanEncode(nodes []HuTNode) {
	//生成Huffman Tree
	n := len(nodes)
	//先权重排序
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Weight < nodes[j].Weight
	})
	nodes0 := make([]HuTNode, n-1)
	nodes = append(nodes, nodes0...)
	//遍历生成非根节点
	for i := 0; i < n-1; i++ {
		if nodes[i+1].Weight <= nodes[n+i-1].Weight {
			nodes[n+i].Weight = nodes[i].Weight + nodes[i+1].Weight
			nodes[i].Parent = &nodes[n+i]
			nodes[i+1].Parent = &nodes[n+i]
			nodes[n+i].Lchild = &nodes[i]
			nodes[n+i].Rchild = &nodes[i+1]
			i++
		} else {
			nodes[n+i].Weight = nodes[i].Weight + nodes[n+i-1].Weight
			nodes[i].Parent = &nodes[n+i]
			nodes[i+n-1].Parent = &nodes[n+i]
			nodes[n+i].Lchild = &nodes[i]
			nodes[n+i].Rchild = &nodes[i+n-1]

		}
	}
	//生成根节点
	nodes[len(nodes)-1].Lchild = &nodes[len(nodes)-2]
	nodes[len(nodes)-1].Rchild = &nodes[len(nodes)-3]
	nodes[len(nodes)-2].Parent = &nodes[len(nodes)-1]
	nodes[len(nodes)-3].Parent = &nodes[len(nodes)-1]
	//记得初始化
	HTCodes = make(map[interface{}]string)
	for i := 0; i <= n-1; i++ {
		code := ""
		node := &nodes[i]
		for true {
			if node.Parent == nil {
				HTCodes[nodes[i].Value] = code
				break
			}

			if node.Parent.Lchild == node {
				code = "0" + code
				node = node.Parent
				continue
			}

			if node.Parent.Rchild == node {
				code = "1" + code
				node = node.Parent
				continue
			}

		}

	}
	fmt.Println(HTCodes)

}

func HuffmanDecode() {

}
