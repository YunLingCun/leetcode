package question

/**

 */
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	ancestor := lowestCommonAncestor(root.Left, p, q)
	if ancestor != nil {
		return ancestor
	}
	ancestor = lowestCommonAncestor(root.Right, p, q)
	if ancestor != nil {
		return ancestor
	}
	var lCount = contain(root.Left, p, q)
	if lCount == 2 {
		return root.Left
	}
	var rCount = contain(root.Right, p, q)
	if rCount == 2 {
		return root.Right
	}
	var count = contain(root, p, q)
	if count == 2 {
		return root
	}
	return nil
}

func contain(root, p, q *TreeNode) int {
	if root == nil || p == nil || q == nil {
		return 0
	}
	var count = 0
	count += contain(root.Left, p, q)
	if count == 2 {
		return count
	}
	count += contain(root.Right, p, q)
	if count == 2 {
		return count
	}
	if root.Val == p.Val || root.Val == q.Val {
		count += 1
	}
	return count
}
