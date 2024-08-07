package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if p.Val == q.Val {
		left := isSameTree(p.Left, q.Left)
		if !left {
			return false
		}
		right := isSameTree(p.Right, q.Right)
		return left == right
	}

	return false
}
