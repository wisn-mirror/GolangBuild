package main

import (
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("README.md")
	if e!=nil{
		panic(e)
	}
	fmt.Println( file)
}
