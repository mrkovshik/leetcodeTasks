package main

import "fmt"

func intToRoman(num int) string {
	var res string
	intRomDict := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "XC",
		100:  "C",
		200:  "CC",
		300:  "CCC",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}
	if num >= 1000 {
		thnd := num / 1000
		res = intRomDict[thnd*1000]
		num = num - thnd*1000
	}
	if num >= 100 {
		hndrd := num / 100
		res = res + intRomDict[hndrd*100]
		num = num - hndrd*100
	}
	if num >= 10 {
		dcm := num / 10
		res = res + intRomDict[dcm*10]
		num = num - dcm*10
	}
	res = res + intRomDict[num]

	return res
}

func main() {
	num := 1994
	rom := intToRoman(num)

	fmt.Printf("Arabic %v is Roman %v", num, rom)
}