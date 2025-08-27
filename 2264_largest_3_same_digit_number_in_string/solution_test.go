package main

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		num string
		want string
	}{
		{
			name: "4 test",
			num: "2300019",
			want: "000",
		},
		{
			name: "5 test",
			num: "014455",
			want: "",
		},
		{
			name: "1 test",
			num: "111",
			want: "111",
		},
		{
			name: "2 test",
			num: "6777133339",
			want: "777",
		},
		{
			name: "3 test",
			num: "42352338",
			want: "",
		},


	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := largestGoodInteger(test.num); got != test.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, test.want)
			}
		})
	}
}
