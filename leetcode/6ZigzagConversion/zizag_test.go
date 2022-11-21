package main

import (
	"fmt"
	"testing"
	"reflect"
)

func TestZigzag(t *testing.T) {
	var tests = []struct {
		s   string
		numRows int
		want    string
	}{
		{"abcdef", 2 ,"acebdf"},
		{"abcdef", 1 ,"abcdef"},
		{"abc,def,ghi", 3 ,"adgb,e,hcfi"},

	}

	for _, tt := range tests {
		testname := fmt.Sprintf("string = %v rowsNum = %v", tt.s, tt.numRows)
		t.Run(testname, func(t *testing.T) {
			ans := convert(tt.s, tt.numRows)
			if reflect.DeepEqual(ans,tt.want) == false {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
