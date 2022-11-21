package main

import (
	"fmt"
	"testing"
	"reflect"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		s   string
		want    string
	}{
		{"aba", "aba"},
		{"ab", "a"},
		{"abadfhdh", "aba"},
		{"1porop1", "1porop1"},
		{"a", "a"},
		{"1porop11porop", "porop11porop"},
		{"1porop11poropaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		
		
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("string = %v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := longestPalindrome(tt.s)
			if reflect.DeepEqual(ans,tt.want) == false {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
