package main
 	

import (
    "fmt"
    "testing"
)

func TestValidPalindrome(t *testing.T) {
    var tests = []struct {
        s string
        want bool
    }{
        {"ede", true},
        {"tgr", false},
		{"e", true},
		{"esd d@ s, e", true},
		{"esd :fd s &^%$&^e", true},
		{"Weew", true},
		{"", true},
		{":>&^", true},
		{"0P", false},
		{",,,,,,,,,,,,acva", false},
		{"Nella's simple hymn: \"I attain my help, Miss Allen.\"", true},
		{"\"Stop!\" nine myriad murmur. \"Put up rum, rum, dairymen, in pots.\"", true},
    }	

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.s)
        t.Run(testname, func(t *testing.T) {
            ans := isPalindrome(tt.s)
            if ans != tt.want {
                t.Errorf("got %t, want %t", ans, tt.want)
            }
        })
    }
}
