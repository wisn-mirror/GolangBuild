package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:="yes 你好golang"//utf-8 可变长 :79 65 73 20 E4 BD A0 E5 A5 BD 67 6F 6C 61 6E 67
	// 0,79  1,65  2,73  3,20  4,4F60  7,597D  10,67  11,6F  12,6C  13,61  14,6E  15,67
	for _,b:=range []byte(s){
		//fmt.Printf("(%d %X)",i,b)
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i,b:=range s{//b就是int32 rune
		fmt.Printf(" %d,%X ",i,b)
	}
	fmt.Println()
	count:=utf8.RuneCountInString(s)
	fmt.Println("Rune Count",count)

	bytes:=[]byte(s)
	for len(bytes)>0{
		char,size:=utf8.DecodeRune(bytes)
		bytes=bytes[size:]
		fmt.Printf("%c",char)
	}
	fmt.Println()

	for i,b:=range []rune(s){
		fmt.Printf(" %d %c ",i,b)
	}


}
