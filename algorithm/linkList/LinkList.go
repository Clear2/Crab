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
}

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		 return nil
	}
	var prev *ListNode
	cur := head
	//for cur != nil {
		// 保存下一个节点
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
		cur.Display()
	//}
	return nil
}

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
	//a.Display()
	reverseList(&a)

	//r := midOrUpMiddleNode(&a)
	//r.Display()
}
