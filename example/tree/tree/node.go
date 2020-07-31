package tree

import "fmt"

type Node struct {
	Value               int
	LeftNode, RightNode *Node
}

func (node Node) PrintValue() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(v int) {
	node.Value = v
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}
