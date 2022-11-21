package main

import (
	"fmt"
	"testing"
	"reflect"
)

func TestGroupAnagrams(t *testing.T) {
	var tests = []struct {
		strs       []string
				want    [][]string
	}{
		{[]string{"fsd", "dsf", "dfsg", "dfs"},[][]string{{"fsd", "dsf","dfs"},{"dfsg"}}},
		{[]string{"f", "f", "", "dfs"},[][]string{{"f", "f"},{""},{"dfs"}}},
		{[]string{"f"},[][]string{{"f"}}},
		{[]string{""},[][]string{{""}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("strings = %v", tt.strs)
		t.Run(testname, func(t *testing.T) {
			ans := groupAnagrams(tt.strs)
			if reflect.DeepEqual(ans,tt.want) == false {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
