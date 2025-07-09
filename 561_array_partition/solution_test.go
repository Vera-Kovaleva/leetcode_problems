package main

import "testing"

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 test",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "1 test",
			args: args{
				nums: []int{6,2,6,5,1,2},
			},
			want: 9,
		},
		{
			name: "1 test",
			args: args{
				nums: []int{6,2,6,5,1,2,1,1},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
