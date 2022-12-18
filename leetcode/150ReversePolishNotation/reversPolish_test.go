package main

import "testing"

type args struct {
	tokens []string
}

func TestPolish(t *testing.T) {

var tests = [] struct {
	name string
	args args
	want int
}{
	{"first", args {[] string {"2","1","+","3","*"}}, 9},
	{"second", args {[] string { "4","13","5","/","+"}}, 6},
	{"third", args {[] string { "10","6","9","3","+","-11","*","/","*","17","+","5","+"}}, 22},
}
	for _,tt:=range tests {
t.Run(tt.name, func(t *testing.T) {
	res:=evalRPN(tt.args.tokens)
	if res != tt.want {
		t.Errorf("got %v, want %v", res, tt.want)
	}
})
	}
}