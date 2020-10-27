package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	array := []*TreeNode{root}
	result := []int{}
	current := root
	for len(array) > 0 {
		current = array[len(array)-1]
		result = append(result, current.Val)
		if len(array) > 1 {
			array = array[0 : len(array)-1]
		} else {
			array = []*TreeNode{}
		}
		if current.Right != nil {
			array = append(array, current.Right)
		}
		if current.Left != nil {
			array = append(array, current.Left)
		}
	}
	return result
}
