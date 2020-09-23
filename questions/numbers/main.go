package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
给出两个非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0开头。

*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	stack := make([][]*TreeNode, 0)
	stack = append(stack, []*TreeNode{t1, t2})
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t[0] == nil || t[1] == nil {
			continue
		}
		t[0].Val += t[1].Val
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			stack = append(stack, []*TreeNode{t[0].Left, t[1].Left})
		}
		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			stack = append(stack, []*TreeNode{t[0].Right, t[1].Right})
		}
	}
	return t1
}

func medianSortedArrays(nums []int) []int {
	var length = len(nums)
	if length == 0 {
		return []int{}
	}
	if length%2 == 0 {
		return []int{length/2 - 1, length / 2}
	} else {
		return []int{length/2 - 1}
	}
}

func findMaxIDX(nums []int, num int) {

}

func findMinIDX(nums []int, num int) {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var medianNums []int
	medianNums = append(medianNums, medianSortedArrays(nums1)...)
	medianNums = append(medianNums, medianSortedArrays(nums2)...)
	if len(medianNums) == 0 {
		return 0
	}

	return 0
}
