package palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				str: "abcba",
			},
			want: true,
		},
		{
			name: "not-ok",
			args: args{
				str: "abcbab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single",
			args: args{
				str: "ababcbad",
			},
			want: "abcba",
		},
		{
			name: "double",
			args: args{
				str: "dabccbabbd",
			},
			want: "abccba",
		},
		{
			name: "double2",
			args: args{
				str: "bb",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.str); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}