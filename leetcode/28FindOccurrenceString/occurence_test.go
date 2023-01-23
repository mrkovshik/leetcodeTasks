package main

import "testing"

type args struct {
	heystack string
	needle   string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test_1", args{"kaplumbaga", "plu"}, 2},
}

func TestOcc(t *testing.T) {

	for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
	res:=strStr(tt.args.heystack,tt.args.needle)
	if res!=tt.want{
		t.Errorf("want %v, got %v", tt.want, res)
	}
})
	}
}