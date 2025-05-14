package tree

type TreeNode struct {
	// The values should be private as we would not expose this to the client directly..
	val   int
	left  *TreeNode
	right *TreeNode
}


