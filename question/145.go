package question

func postOrderTraversal(root *TreeNode) []int {
	type firstNode struct {
		node  *TreeNode
		First bool
	}
	var stack = make([]*firstNode, 0, 0)
	if root != nil {
		stack = append(stack, &firstNode{
			node:  root,
			First: true,
		})
	} else {
		return nil
	}
	var postOrder = make([]int, 0, 0)
	for len(stack) > 0 {
		now := stack[len(stack)-1]
		if now.First {
			if now.node.Right != nil {
				stack = append(stack, &firstNode{
					node:  now.node.Right,
					First: true,
				})
			}
			if now.node.Left != nil {
				stack = append(stack, &firstNode{
					node:  now.node.Left,
					First: true,
				})
			}
			now.First = false
		} else {
			postOrder = append(postOrder,now.node.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return postOrder
}

func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

