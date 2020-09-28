package question

func connect(root *Node) *Node {
	//var stack = make(map[int]*Node)
	//Connect(root, stack, 0)
	Connect2(root)
	return root
}

func Connect(root *Node, stack map[int]*Node, level int) {
	if root == nil {
		return
	}
	Connect(root.Left, stack, level+1)
	Connect(root.Right, stack, level+1)
	if node, ok := stack[level]; ok {
		node.Next = root
	}
	stack[level] = root
}

func Connect2(root *Node) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Left != nil && root.Right == nil {
		root.Left.Next = GetNext(root.Next)
	}
	if root.Right != nil {
		root.Right.Next = GetNext(root.Next)
	}
	Connect2(root.Right)
	Connect2(root.Left)
}

func GetNext(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return node.Left
	}
	if node.Right != nil {
		return node.Right
	}
	if node.Next != nil {
		return GetNext(node.Next)
	}
	return nil
}
