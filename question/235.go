package question

/**
	1. root 等于 p or q 只需要查看子节点中是否有另外一个
    2. root 不等于 p or q
		2.1 root 的子节点中包含其中一个
        2.2 root 的子节点中不包含
        2.3 root 子节点中包含二者
*/

var ancestor *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor = nil
	return LowestCommonAncestor(root, p, q)
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if LowestCommonAncestor(root.Left, p, q) != nil {
		return ancestor
	}
	if LowestCommonAncestor(root.Right, p, q) != nil {
		return ancestor
	}
	if LowestCommonAncestor(root, p, q) != nil {
		return ancestor
	}
	if ancestor != nil {
		return ancestor
	} else {
		if contain(root.Right, p) && contain(root.Right, q) {
			ancestor = root.Right
			return ancestor
		}
		if contain(root.Left, p) && contain(root.Left, q) {
			ancestor = root.Left
			return ancestor
		}
		if contain(root.Left, p) && contain(root.Right, q) {
			ancestor = root
			return ancestor
		}
		if contain(root.Right, p) && contain(root.Left, q) {
			ancestor = root
			return ancestor
		}
	}
	return nil
}

func contain(root, n *TreeNode) bool {
	if n == nil {
		return false
	}
	if root == nil {
		return false
	}
	if root.Val == n.Val {
		return true
	}
	if contain(root.Left, n) {
		return true
	}
	if contain(root.Right, n) {
		return true
	}
	return false
}
