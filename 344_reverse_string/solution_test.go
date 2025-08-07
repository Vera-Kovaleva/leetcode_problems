package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "test 1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			require.Equal(t, tt.want, tt.s)
		})
	}
}
