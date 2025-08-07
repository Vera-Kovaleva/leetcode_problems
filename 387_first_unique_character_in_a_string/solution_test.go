package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "test 1",
			s:    "itwqbtcdprfsuprkrjkausiterybzncbmdvkgljxuekizvaivszowqtmrttiihervpncztuoljftlxybpgwnjb",
			want: 61,
		},
		{
			name: "test 2",
			s:    "loveleetcode",
			want: 2,
		},
		{
			name: "test 3",
			s:    "aabb",
			want: -1,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
