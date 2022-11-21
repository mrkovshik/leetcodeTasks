package main
 	

import (
    "fmt"
    "testing"
)

func TestSearchRange(t *testing.T) {
    var tests = []struct {
        nums []int
		target int
        want []int
    }{
        {[]int {1,2,3,3,4}, 3, [] int {2,3}},
		{[]int {1,2,3,3,3}, 3, [] int {2,4}},
		{[]int {1,2,3,1,1}, 3, [] int {2,2}},
		{[]int {1,2,3,1,1}, 7, [] int {-1,-1}},
		{[]int {1,2,3,3,4}, 1, [] int {0,0}},
    }	


    for _, tt := range tests {
        testname := fmt.Sprintf("%d%d", tt.nums, tt.target)
        t.Run(testname, func(t *testing.T) {
            ans := searchRange(tt.nums, tt.target)
            if ans[0]  != tt.want[0]||tt.want[1]!=ans[1] {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
