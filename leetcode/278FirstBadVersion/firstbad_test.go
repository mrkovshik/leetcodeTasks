package main
 	

import (
    "fmt"
    "testing"
)

func TestFindFirstBad(t *testing.T) {
    var tests = []struct {
        rng int
		want int
    }{
        {40, 3},
		{1, 1},
		{2, 1},
		{2, 2},
    }	


    for _, tt := range tests {
        testname := fmt.Sprintf("%d", tt.rng)
        t.Run(testname, func(t *testing.T) {
            ans := firstBadVersion(tt.rng)
            if ans  != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
