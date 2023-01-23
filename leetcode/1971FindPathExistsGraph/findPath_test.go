package main

import (
	"testing"
)

type args struct {
	n           int
	edges       [][]int
	source      int
	destination int
}

func TestFindPath(t *testing.T) {
var tests = [] struct{
	name string
	args args
	want bool
}{
	{"first", args {3, [][]int {{0, 1}, {1, 2}, {2, 0},}, 0,2}, true},
	{"sec", args {6, [][]int {{0, 1}, {0, 2}, {3, 5},{5, 4},{4, 3}}, 0,5}, false},
	{"third", args {10, [][]int {{4, 3}, {1, 4}, {4, 8},{1, 7},{6,4},{4,2},{7,4}, {4,0},   {0,9},{5,4}}, 5,9}, true},
}
  for _,tt:= range tests {
	t.Run(tt.name, func(t *testing.T) {
		res:= validPath (tt.args.n, tt.args.edges, tt.args.source, tt.args.destination)
		if res!=tt.want{
			t.Errorf("want %v, got %v", tt.want, res)
		}
	})
  }
}