package main

import "testing"

func Test_isUniqString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"same",
			args{"abcdefg"},
			true,
		},
		{
			"same",
			args{"abcdefa"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUniqString(tt.args.str); got != tt.want {
				t.Errorf("isUniqString() = %v, want %v", got, tt.want)
			}
		})
	}
}
