package addTwoNumber

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"example 1",
			args{generateListNode([]int{2, 4, 3}), generateListNode([]int{5, 6, 4})},
			generateListNode([]int{7, 0, 8}),
		},
		{
			"example 2",
			args{generateListNode([]int{0}), generateListNode([]int{0})},
			generateListNode([]int{0}),
		},
		{
			"example 3",
			args{generateListNode([]int{9, 9, 9, 9, 9, 9, 9}), generateListNode([]int{9, 9, 9, 9})},
			generateListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateListNode(input []int) *ListNode {
	result := &ListNode{}
	cur := result
	for i := range input {
		cur.Val = input[i]
		if i != len(input)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return result
}
