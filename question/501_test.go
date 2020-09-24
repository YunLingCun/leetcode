package question

import "testing"

func TestFindModeFromSortedArrays(t *testing.T) {
	arrays := []int{-1,-1,-1,-1,-1,-1,0,0,0, 1,1,1, 2, 3, 4, 5, 5, 6, 6, 6,6,6,6, 7,7,7, 7, 8, 8, 8, 8, 9, 9}
	t.Logf("%v", FindModeFromSortedArrays(arrays))
	t.Logf("%v", FindModeFromSortedArrays([]int{1,2,3}))
	t.Logf("%v", FindModeFromSortedArrays([]int{1,1,2,3,3,4,5,5}))
	t.Logf("%v", findMode(&TreeNode{
		Val:   2147483647,
		Left:  nil,
		Right: nil,
	}))
}
