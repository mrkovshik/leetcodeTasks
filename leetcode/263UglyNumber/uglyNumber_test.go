package main
 	

import (
    "fmt"
    "testing"
)

func TestIsUgly(t *testing.T) {
    var tests = []struct {
        n int
        want bool
    }{
        {1, true},
		{6, true},
		{14, false},
		{21, false},
		{23, false},
		{25, true},
		{8, true},
		{12, true},
		{26, false},

       
    }	

    for _, tt := range tests {
        testname := fmt.Sprintf("%v", tt.n)
        t.Run(testname, func(t *testing.T) {
            ans := isUgly(tt.n)
            if ans != tt.want {
                t.Errorf("got %t, want %t", ans, tt.want)
            }
        })
    }
}
