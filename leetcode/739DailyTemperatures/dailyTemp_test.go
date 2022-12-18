package main

import (
	"reflect"
	"testing"
)
 type args struct {
temperatures [] int
 }
func TestDailyTemp(t *testing.T) {
tests:= []struct {
	name string
	args args
	want []int
}{
	{"first",args {[]int{73, 74, 75, 71, 69, 72, 76, 73}}, [] int {1,1,4,2,1,1,0,0}},
	{"first",args {[]int{30,40,50,60}}, [] int {1,1,1,0}},
	{"first",args {[]int{30,60,90}}, [] int {1,1,0}},
	{"first",args {[]int{30}}, [] int {0}},
	{"first",args {[]int{89,62,70,58,47,47,46,76,100,70}}, [] int {8,1,5,4,3,2,1,1,0,0}},
	{"first",args {[]int{34,80,80,34,34,80,80,80,80,34}}, [] int {1,0,0,2,1,0,0,0,0,0}},	
}
for _,tt:= range tests {
	t.Run(tt.name, func(t *testing.T){
		res:= dailyTemperatures(tt.args.temperatures)
		if !reflect.DeepEqual(res,tt.want){
			t.Errorf("got %d, want %d", res, tt.want)
		}
	})
}
	
	
}