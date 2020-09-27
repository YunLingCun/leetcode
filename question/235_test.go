package question

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	var root = new(TreeNode)
	root.Left = new(TreeNode)
	root.Val = 1
	root.Left.Val = 2
	if ancestor := LowestCommonAncestor(root, root, root.Left); ancestor != nil {
		t.Logf("%+v", ancestor)
	}else {
		t.Errorf("ancestor not found")
	}
}
