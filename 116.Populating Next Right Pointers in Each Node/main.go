package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	leftPtr := root.Left
	rightPtr := root.Right
	for leftPtr != nil {
		leftPtr.Next = rightPtr
		leftPtr = leftPtr.Right
		rightPtr = rightPtr.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

func main() {

}
