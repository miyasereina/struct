package B_

type BTNode struct {
	n      int
	parent *BTNode
	key    []interface{}
	childs []*BTNode
}

func contribute(n int) *BTNode {
	return &BTNode{
		n:      n,
		childs: make([]*BTNode, n+1),
		key:    make([]interface{}, n),
	}
}
