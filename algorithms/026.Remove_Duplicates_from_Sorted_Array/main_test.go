package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			"Example 1",
			args{
				[]int{1, 1, 2},
			},
			2,
			[]int{1, 2},
		},
		{
			"Example 2",
			args{
				[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantNums[:got]) {
				t.Errorf("nums = %v, want %v", tt.args.nums, tt.wantNums)
			}
		})
	}
}
