package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example 1",
			args{
				head: []int{1, 2, 3},
			},
			[]int{3, 2, 1},
		},
		{
			"Example 2",
			args{
				head: []int{1, 2},
			},
			[]int{2, 1},
		},
		{
			"Example 3",
			args{
				head: []int{},
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(sliceToList(tt.args.head)); !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}

func sliceToList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	var pre *ListNode

	for i := len(s) - 1; i >= 0; i-- {
		if pre == nil {
			pre = &ListNode{
				Val: s[i],
			}
			continue
		}
		pre = &ListNode{
			Val:  s[i],
			Next: pre,
		}
	}

	return pre
}

func listToSlice(n *ListNode) []int {
	s := []int{}
	for n != nil {
		s = append(s, n.Val)
		n = n.Next
	}

	return s
}
