package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := invertTree(root.Left)
	left := invertTree(root.Right)
	root.Right = right
	root.Left = left
	return root
}
