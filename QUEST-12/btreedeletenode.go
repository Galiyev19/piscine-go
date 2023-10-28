package piscine

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		// Node to be deleted found

		// Case 1: Node with only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 3: Node with two children
		// Find the in-order successor (smallest node in the right subtree)
		root.Data = findMin(root.Right).Data
		// Delete the in-order successor
		root.Right = BTreeDeleteNode(root.Right, &TreeNode{Data: root.Data})
	}

	return root
}
