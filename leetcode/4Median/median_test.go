package main
 	

import (
    "fmt"
    "testing"
)

func TestMedian(t *testing.T) {
    var tests = []struct {
        nums1 []int
		nums2 []int
        want float64
    }{
        {[] int{2,2,3}, [] int{1,5,7},2.5},
		{[] int{2}, [] int{1,5,7},3.5},
		{[] int{2,8}, [] int{1,5,7},5},
        {[] int{1,3}, [] int{2},2},
   
    }	

    for _, tt := range tests {
        testname := fmt.Sprintf("nums1= %v, nums2 = %v", tt.nums1, tt.nums2)
        t.Run(testname, func(t *testing.T) {
            ans := findMedianortedArrays(tt.nums1,tt.nums2)
            if ans != tt.want {
                t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}
