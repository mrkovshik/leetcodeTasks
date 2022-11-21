package main

import "fmt"

func longestPalindrome(s string) string {
	var res string
	res=string(s[0])
	max:=res
	for i:=0; i<len(s); i++ {
		for j:=i+1; j<len(s); j++{
			if s[i]== s[j]{
				leng:=(j-i)+1
				pal:=make ([] byte, leng)
				pal[0]=s[i]
				pal[leng-1]=s[j]
				if j-i==1{
					res=string(pal)
				} else if j-i==2 {
					pal[1]=s[i+1]
					res=string(pal)
				} else {
			for k:=1; k<j-i; k++ {
					if s[i+k]==s[j-k]{
						switch{
					case (j-k)-(i+k)==1:
						pal[k]=s[i+k]
						pal[leng-1-k]=s[j-k]
						res=string(pal)
						break
					case (j-k)-(i+k)==2:
						pal[k]=s[i+k]
						pal[leng-1-k]=s[j-k]
						pal[k+1]=s[i+k+1]
						res=string(pal)
						break
					default:
						pal[k]=s[i+k]
						pal[leng-1-k]=s[j-k]
					} 
					} else {
						res=string(s[0])
						break
					}
				}
			}
				if len(res)> len(max) {
					max=res
				}
			} 
		}
	}
return max
}

func main() {
	s := "aba"
	fmt.Printf("The longest palindrome in %v is %v", s, longestPalindrome(s))

}