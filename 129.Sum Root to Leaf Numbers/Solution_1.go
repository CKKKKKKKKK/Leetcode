/*
 *	Solution_1: dfs algorithm
 */
package sumNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, parent int, sum *int) {
	current := parent*10 + root.Val
	if root.Left != nil && root.Right != nil {
		dfs(root.Left, current, sum)
		dfs(root.Right, current, sum)
	} else if root.Left != nil {
		dfs(root.Left, current, sum)
	} else if root.Right != nil {
		dfs(root.Right, current, sum)
	} else {
		*sum = *sum + current
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := 0
	sum := &a
	dfs(root, 0, sum)
	return *sum
}
