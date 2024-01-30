package main

// https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return n
}

type ListNode struct {
	Val  int
	Next *ListNode
}
