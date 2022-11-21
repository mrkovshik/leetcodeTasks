package main

import (
	"fmt"
	"testing"
	"reflect"
)

func TestReverseString(t *testing.T) {
	var tests = []struct {
		s       string
				want    string
	}{
		{"Roman narkoman", "narkoman Roman"},
		{"   Roman   narkoman   ", "narkoman Roman"},
		{"narkoman   ", "narkoman"},
		{"narkoman", "narkoman"},
		{"Roman   narkoman  sdfgasfg   afdgasfg ", "afdgasfg sdfgasfg narkoman Roman"},
		
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("string = %v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := reverseWords(tt.s)
			if reflect.DeepEqual(ans,tt.want) == false {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
