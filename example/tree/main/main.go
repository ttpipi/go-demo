package main

import (
	"demo/example/tree/tree"
	"fmt"
)

func main() {
	//构造
	root := tree.Node{Value: 1, LeftNode: nil, RightNode: nil}
	root.LeftNode = &tree.Node{Value: 2}
	root.RightNode = &tree.Node{Value: 3, LeftNode: nil, RightNode: nil}
	root.LeftNode.RightNode = tree.CreateNode(4)
	root.RightNode.LeftNode = new(tree.Node)

	root.RightNode.LeftNode.SetValue(5)

	root.Traverse()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	maxNode := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if maxNode < node.Value {
			maxNode = node.Value
		}
	}
	fmt.Println("Max Node:", maxNode)
}
