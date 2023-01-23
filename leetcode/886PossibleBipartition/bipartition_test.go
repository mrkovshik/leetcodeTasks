package main

import "testing"

type args struct {
	n        int
	dislikes [][]int
}

var tests = []struct {
	name string
	args args
	want  bool
}{
	{"test_1", args{3, [][]int{{1, 2}, {1, 3}, {3, 2}}}, false},
	{"test_2", args{4, [][]int{{1, 2}, {1, 3}, {4, 2}}}, true},
	{"test_3", args{5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}}}, false},
	{"test_4", args{1, [][]int{}}, true},
}

func TestBipart(t *testing.T) {
for _,tt:= range tests {
	t.Run(tt.name,func(t *testing.T) {
		res:=possibleBipartition(tt.args.n, tt.args.dislikes)
		if res!=tt.want {
t.Errorf("want= %v, got= %v", tt.want,res)
		}
	})
}
}