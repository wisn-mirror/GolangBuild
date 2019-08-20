package main

import (
	"fmt"
	"os"
)

func findLong(str string) (int, string) {
	lastcode:=make(map[rune]int)
	start:=0
	maxlength:=0
	for i,value:=range []rune(str){
		lastindex,ok:=lastcode[value]
		if ok{
			if lastindex>=start{
				start=i
			}
		}
		if i-start>=maxlength{
			maxlength=i-start+1
		}
		lastcode[value]=i
		//fmt.Println(i,value)
	}
	//fmt.Println(lastcode)
	return maxlength,"str";
}


func main() {
	args := os.Args
	fmt.Println(findLong("abccdefff"))
	fmt.Println(findLong("中国人人"))
	fmt.Println(findLong(""))
	fmt.Println(findLong("a"))
	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
}
