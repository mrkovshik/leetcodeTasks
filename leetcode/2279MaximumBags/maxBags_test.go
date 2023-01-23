package main

import "testing"

type args struct {
	capacity        []int
	rocks           []int
	additionalRocks int
}

func TestMaxBags(t *testing.T) {
var tests = [] struct{
	name string
	args args
	want int
}{
{"test_1", args{[]int{2,3,4,5}, [] int {1,2,4,4}, 2}, 3},
{"test_2", args{[]int{10,2,2}, [] int {2,2,0}, 90}, 3},
{"test_3", args{[]int{0,0,0}, [] int {0,0,0}, 0}, 3},
{"test_4", args{[]int{54,18,91,49,51,45,58,54,47,91,90,20,85,20,90,49,10,84,59,29,40,9,100,1,64,71,30,46,91}, [] int {14,13,16,44,8,20,51,15,46,76,51,20,77,13,14,35,6,34,34,13,3,8,1,1,61,5,2,15,18}, 77}, 13},
}
for _, tt:= range tests {
	t.Run(tt.name,func(t *testing.T) {
		result:=maximumBags(tt.args.capacity,tt.args.rocks,tt.args.additionalRocks)
		if result!=tt.want {
			t.Errorf("got %v, want %v", result, tt.want)
		}
	})
}
}