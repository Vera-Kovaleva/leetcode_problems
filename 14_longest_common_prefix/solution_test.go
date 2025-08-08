package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "test 3",
			strs: []string{"ab", "a"},
			want: "a",
		},
		{
			name: "test 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "test 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
