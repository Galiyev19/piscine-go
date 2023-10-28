package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	Left := BTreeLevelCount(root.Left)
	Right := BTreeLevelCount(root.Right)

	return Max(Left, Right) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
