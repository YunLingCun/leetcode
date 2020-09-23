package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CameraCover(root *TreeNode) (int, int) {
	var rightMin = 0
	var leftMin = 0
	var rightVal = -1
	var leftVal = -1
	if root.Left != nil {
		leftMin, leftVal = CameraCover(root.Left)
	}
	if root.Right != nil {
		rightMin, rightVal = CameraCover(root.Right)
	}
	if leftVal == 0 || rightVal == 0 {
		return rightMin + leftMin + 1, 2
	}
	if leftVal == 2 || rightVal == 2 {
		return rightMin + leftMin, 1
	}
	return rightMin + leftMin, 0
}

func MinCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var min, val = CameraCover(root)
	if val == 0 {
		return min + 1
	} else {
		return min
	}
}
