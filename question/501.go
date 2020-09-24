package question

/**
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

1. 在一个有序数组中找到众数
2. 将二叉搜索树中序遍历
*/

func FindModeFromSortedArrays(nums []int) []int {
	var lastCount int
	var lastNum int
	var baseCount int
	var modes []int
	for idx, num := range nums {
		if idx == 0 {
			lastCount = 1
			lastNum = num
			baseCount = 1
			modes = append(modes, num)
			continue
		}
		if num == lastNum {
			lastCount += 1
		} else {
			lastCount = 1
			lastNum = num
		}
		if lastCount == baseCount {
			modes = append(modes, num)
		} else if lastCount > baseCount {
			modes = []int{num}
			baseCount = lastCount
		}
	}
	return modes
}

type dataResult struct {
	LastCount int
	LastNum   int
	BaseCount int
	Modes     []int
	NotFirst  bool
}

var result dataResult

func findModeInOrderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	findModeInOrderTraversalRecursion(root.Left)
	if !result.NotFirst {
		result.LastCount = 1
		result.LastNum = root.Val
		result.BaseCount = 1
		result.Modes = append(result.Modes, root.Val)
		result.NotFirst = true
	} else {
		if root.Val == result.LastNum {
			result.LastCount += 1
		} else {
			result.LastCount = 1
			result.LastNum = root.Val
		}
		if result.LastCount == result.BaseCount {
			result.Modes = append(result.Modes, root.Val)
		} else if result.LastCount > result.BaseCount {
			result.Modes = []int{root.Val}
			result.BaseCount = result.LastCount
		}
	}
	findModeInOrderTraversalRecursion(root.Right)
}

func InOrderTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	nums = append(nums, InOrderTraversalRecursion(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, InOrderTraversalRecursion(root.Right)...)
	return nums
}

func InOrderTraversalNonRecursive(root *TreeNode) []int {
	var stack []*TreeNode
	var p = root
	var nums []int
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			nums = append(nums, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return nums
}

func findModeByInOrderTraversalNonRecursive(root *TreeNode) []int {
	var stack []*TreeNode
	var p = root
	var lastCount int
	var lastNum int
	var baseCount int
	var modes []int
	var first = true
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			var val = p.Val
			stack = stack[:len(stack)-1]
			p = p.Right
			if first {
				lastCount = 1
				lastNum = val
				baseCount = 1
				modes = append(modes, val)
				first = false
				continue
			}
			if val == lastNum {
				lastCount += 1
			} else {
				lastCount = 1
				lastNum = val
			}
			if lastCount == baseCount {
				modes = append(modes, val)
			} else if lastCount > baseCount {
				modes = []int{val}
				baseCount = lastCount
			}
		}
	}
	return modes
}

func findMode(root *TreeNode) []int {
	findModeInOrderTraversalRecursion(root)
	return result.Modes
}
