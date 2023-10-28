package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}

	leftres := BTreeSearchItem(root.Left, elem)

	if leftres != nil {
		return leftres
	}

	if root.Data == elem {
		return root
	}

	rightRes := BTreeSearchItem(root.Right, elem)

	if rightRes != nil {
		return rightRes
	}

	return nil
}
