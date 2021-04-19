package main

import "fmt"

// 输入链表头节点，奇数长度返回中点，偶数长度返回上中点
// 输入链表头节点，奇数长度返回中点，偶数长度返回下中点
// 输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
// 输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Display() {
	for l != nil {
		fmt.Printf("%v->", l.Val)
		l = l.Next
	}
	fmt.Printf("\n")
}
/*
 * 打印两个有序链表的公共部分
 * 谁小谁移动，相等了打印，打印完成共同移动，直到一个结束为止
 */


// 反转链表 1->2 ->3 ->3
// pre当前节点的前一个指针 ,cur记录当前节点  tmp保存当前节点的后续节点
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		 return nil
	}
	var cur *ListNode
	prev := head
	//for i:=0; i<1;i++ {
		// 保存下一个节点
		//temp := cur.Next
		prev.Next = nil
		prev.Display()
		cur = prev
		cur.Display()
		//prev.Display()
		//prev = cur
		//prev.Display()
		//cur = temp

	//}
	//prev = head
	return nil
}

/**
	temp := cur.Next
	cur.Next = prev
	prev = cur
	cur = temp
	cur.Display()
 */
func midOrUpMiddleNode(head *ListNode) *ListNode {
	fmt.Println(head.Next == nil)
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow
}


// 给定一个单链表的头节点head, 请判断该链表是否为回文结构
// 笔试 遍历推进栈
// 快慢指针 ->

func main() {
	a := ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 3}
	a.Next.Next.Next = &ListNode{Val: 4}
	a.Next.Next.Next.Next = &ListNode{Val: 5}
	reverseList(&a)
	//a.Display()

	//r := midOrUpMiddleNode(&a)
	//r.Display()
}
