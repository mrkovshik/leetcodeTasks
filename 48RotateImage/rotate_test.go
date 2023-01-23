package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
var tests = [] struct {
	name string
	matrix [][]int
	want [][] int
}{
	{"test_1", [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}}},
	{"test_2", [][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{7,4,1},{8,5,2},{9,6,3}}},
}

for _,tt:=range tests {
	t.Run(tt.name, func(t *testing.T) {
		rotate(tt.matrix)
		if !reflect.DeepEqual(tt.matrix, tt.want){
			t.Errorf("want %v, got %v", tt.want,tt.matrix)
		}
	})
}

}