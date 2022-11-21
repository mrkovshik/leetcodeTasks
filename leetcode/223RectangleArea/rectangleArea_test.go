package main
 	

import (
    "fmt"
    "testing"
)

func TestRectArea(t *testing.T) {
    var tests = []struct {
    ax1 int
	ay1 int
	ax2 int
	ay2 int
	bx1 int
	by1 int
	bx2 int
	by2 int
	want int
    }{
        {-3,0,3,4,0,-1,	9,2, 45},
		{-5,-1,5,6,-2,-5,8,2, 119},
		{-2,-2,2,2,3,3,5,5, 20},
		{-2,-2,2,2,-2,-4,2,-2, 24},
		    }	


    for _, tt := range tests {
        testname := fmt.Sprintf("%v%v%v%v%v%v%v%v", tt.ax1, tt.ay1,tt.ax2,tt.ay2, tt.bx1, tt.by1,tt.bx2,tt.by2)
        t.Run(testname, func(t *testing.T) {
            ans := computeArea(tt.ax1, tt.ay1,tt.ax2,tt.ay2, tt.bx1, tt.by1,tt.bx2,tt.by2)
            if ans  != tt.want{
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
