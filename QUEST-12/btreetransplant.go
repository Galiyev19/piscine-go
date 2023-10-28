package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// Check if the node is nil or not in the tree
	if root == nil || node == nil {
		return root
	}

	// If node is the root, update the root with the replacement
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// If node is the left child of its parent
		node.Parent.Left = rplc
	} else {
		// If node is the right child of its parent
		node.Parent.Right = rplc
	}

	// If the replacement node is not nil, update its parent
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
