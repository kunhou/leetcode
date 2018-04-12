package addTwoNumber

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret

	sum := 0
	for true {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Val = sum % 10
		sum = sum / 10

		if l1 != nil || l2 != nil || sum != 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		} else {
			break
		}
	}
	return ret
}
