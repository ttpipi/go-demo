package tree

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.PrintValue()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.LeftNode.TraverseFunc(f)
	f(node)
	node.RightNode.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
