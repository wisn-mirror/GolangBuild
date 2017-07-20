// Var
package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	//枚举
	m_a = "A"
	m_b
	//从0开始 每定义一个常量就会加1  m_c=2
	m_c = iota
	m_d
)
const (
	J_1 = iota
	J_2
	J_3
	J_4
)

func VarTest() {
	fmt.Println("\n**********************")
	fmt.Println(m_a)
	fmt.Println(m_b)
	fmt.Println(m_c)
	fmt.Println(m_d)
	fmt.Println("**********************")
	fmt.Println("\n********jjjjjjjjjjjjj**************")
	fmt.Println(J_1)
	fmt.Println(J_2)
	fmt.Println(J_3)
	fmt.Println(J_4)
	fmt.Println("**********************")
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.Max(float64(33), float64(5)))
	var a int = 65
	b := string(a)
	var c string = strconv.Itoa(a)
	strconv.Atoi(c)
	var d string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(strconv.Atoi(c))
	fmt.Println("Hello World!")
	const e = 1
	const f int = 33
	const g = "helloworld           "
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(len(g))
}

/*
6:	0110
11:	1011
---------------
&	0010	2
|	1111	15
^	1101	13
&^	0100	4

*/
func TestSuan() {
	fmt.Println("---------TestSuan")
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	//&  和&&
	ts_A := 0
	if ts_A > 0 && (2/ts_A) > 1 {

		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	fmt.Println("uuuuno")
	fmt.Println("tttno")

}
