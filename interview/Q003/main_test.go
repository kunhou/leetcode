package q003

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			"single character",
			args{"a"},
			"a",
		},
		{
			"longstring",
			args{"abcdefgh"},
			"hgfedcba",
		},
		{
			"longstring 2",
			args{"1234567890abcde"},
			"edcba0987654321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := reverseString(tt.args.s); gotOutput != tt.wantOutput {
				t.Errorf("reverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
