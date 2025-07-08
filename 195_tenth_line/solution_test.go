package main

import "testing"

func Test_temLines(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first test",
			args: args{fileName: "1Line.txt"},
			want: "1 line",
		},
		{
			name: "first test",
			args: args{fileName: "2Line.txt"},
			want: "1 line\n2 line",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temLines(tt.args.fileName); got != tt.want {
				t.Errorf("temLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
