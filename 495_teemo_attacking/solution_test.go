package main

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	tests := []struct {
		name string
		timeSeries []int
		duration   int
		want       int
	}{
		{
			name: "test 1",
			timeSeries: []int{1,4},
			duration: 2,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPoisonedDuration(tt.timeSeries, tt.duration)
			// TODO: update the condition below to compare got with tt.want.
			if got!=tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
