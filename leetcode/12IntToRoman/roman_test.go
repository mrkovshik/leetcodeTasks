package main
 	

import (
    "fmt"
    "testing"
)

func TestIntToRoman (t *testing.T){
var tests =[] struct {
	num int
	want string
}{
	{1986, "MCMLXXXVI"},
	{1, "I"},
	{2345, "MMCCCXLV"},
}

for _, tt := range tests {
testname:= fmt.Sprintf("Arabic= %v", tt.num)
t.Run(testname, func(t *testing.T){
res:=intToRoman(tt.num)
if res != tt.want {
	t.Errorf("got %v, want %v", res, tt.want)
}
})
}
}