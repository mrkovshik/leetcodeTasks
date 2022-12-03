package main
 	

import (
    "fmt"
    "testing"
)

func TestReverseInt(t *testing.T) {
    var tests = []struct {
        x int
        want int
    }{
        {123, 321},
		{12300, 321},
		{10203, 30201},
		{0, 0},
		{10, 1},
   
    }	

    for _, tt := range tests {
        testname := fmt.Sprintf("x= %v", tt.x)
        t.Run(testname, func(t *testing.T) {
            ans := reverse(tt.x)
            if ans != tt.want {
                t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}
