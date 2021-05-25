package main

// 二叉查找树 左边的值都比根节点小，右边的根都比根节点大

func isBST(head *TreeNode) bool {
	if head == nil {
		return true
	}
	left := isBST(head.Left)
	if !left {
		return false
	}
	if head.Val <= preValue {
		.
	}
}

func searchBTS(root *TreeNode, val int)  *TreeNode{
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBTS(root.Right, val)
	} else {
		return searchBTS(root.Left, val)
	}
}

