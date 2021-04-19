package main

import (
	"fmt"
)

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

// 前序遍历 头 => 左 => 右
func PreOrderRecur(root * TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	PreOrderRecur(root.Left)
	PreOrderRecur(root.Right)
}
// 中序遍历 左 => 头 => 右
func InOrderRecur(root * TreeNode) {
	if root == nil {
		return
	}
	InOrderRecur(root.Left)
	fmt.Printf("%d", root.Val)
	InOrderRecur (root.Right)
}

// 后序遍历
func PosOrderRecur(root *TreeNode)  {
	if root == nil {
		return
	}
	PosOrderRecur(root.Left)
	PosOrderRecur(root.Right)
	fmt.Printf("%d", root.Val)
}

/**
非递归遍历 先把头节点放到栈中
1.从栈中弹出一个节点current
2.打印 current
3.先右 在左 入栈 如果有的话
4.周而复始
*/
func PreOrderUnRecur(root *TreeNode)  {
	if root != nil {
		stack := []*TreeNode{}
		stack = append(stack, root)
		for len(stack) != 0 {
			root = stack[len(stack) -1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d", root.Val)
			if  root.Right != nil{
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
}

/** 后序遍历
中序:  头 左 右
中序': 头 右 左 1 3 7 6 2 4 5
1.弹出current 放到 收集栈 【】
2.先压左  后压右
 */
func posOrderUnRecur(root *TreeNode)  {
	if root != nil {
		 stack1 := []*TreeNode{}
		 stack2 := []*TreeNode{}
		 stack1 = append(stack1, root)
		for len(stack1) != 0 {
			root = stack1[len(stack1)-1]
			// 弹出节点
			stack1 = stack1[:len(stack1)-1]
			// 压入收集栈
			stack2 = append(stack2, root)

			// 先左后右
			if  root.Left != nil {
				stack1 = append(stack1, root.Left)
			}
			if root.Right != nil {
				stack1 = append(stack1, root.Right)
			}
		}

		// 遍历收集栈
		for len(stack2) != 0 {
			tree := stack2[len(stack2) -1]
			fmt.Printf("%d", tree.Val)
			stack2 = stack2[:len(stack2)-1]

		}
	}
}
/*
每棵子树 整棵树左节点进栈
依次弹出左节点的过程中 打印
对弹出节点的右树 周而复始
1 2 4
*/



/**
判断是否是搜索二叉树，中序遍历，看下是否是
*/

var preValue = 0
func checkBST(head *TreeNode)  bool {
	if head == nil {
		return true;
	}
	// 如果左树不是二叉树，直接返回
	isLeftBST := checkBST(head.Left)
	if isLeftBST {
		return false
	}
	// preValue 是否是递增的
	if head.Val <= preValue {
		return false
	} else {
		preValue = head.Val
	}

	// 判断右树
	return checkBST(head.Right)
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val:4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	/**
	 	1
	2      3
4     5   6  7
	 */
	//PreOrder(&root)
	//InOrder(&root)
	//PosOrderRecur(&root)
	//PreOrderUnRecur(&root)
	//posOrderUnRecur(&root)
}
// 2 6
// 1 3
