package isSubsequence

import (
	"reflect"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// test1
		{
			name: "test1",
			args: args{
				s: "acb",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortStringT(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want map[rune][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortStringT(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortStringT() = %v, want %v", got, tt.want)
			}
		})
	}
}
