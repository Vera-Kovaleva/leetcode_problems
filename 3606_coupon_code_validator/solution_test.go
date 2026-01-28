package main

import "testing"

func isEqual(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i<len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func Test_validateCoupons(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		code         []string
		businessLine []string
		isActive     []bool
		want         []string
	}{
		{
			name:         "test 1",
			code: []string{"SAVE20","","PHARMA5","SAVE@20"},
			businessLine: []string{"restaurant","grocery","pharmacy","restaurant"},
			isActive:     []bool{true,true,true,true},
			want:         []string{"PHARMA5","SAVE20"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateCoupons(tt.code, tt.businessLine, tt.isActive)
			// TODO: update the condition below to compare got with tt.want.
			if !isEqual(tt.want, got) {
				t.Errorf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
