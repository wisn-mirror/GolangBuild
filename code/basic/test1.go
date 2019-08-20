package main

import "fmt"

func main() {
	name := "ddd"
	i, s := test3(&name)
	fmt.Println(i, s)

	 array2:=[]int{2,3}
	fmt.Println(array2)
	array1:=[]int{2,3}
	fmt.Println(array1)
	fmt.Println(array1==nil)
	/**结果：
	[]
	true
	*/
}
func test3(name *string) (int, string) {
	return 1, *name
}

func test2() {
	a, b := TestReturnValue()
	fmt.Println(a, b)
	c, _ := TestReturnValue()
	fmt.Println(c)
}
func TestReturnValue() (int, int) {
	return 10, 20
}

func test1() {
	name1, name2 := "evey", "wisn"
	fmt.Println(name1, name2)
	var a = 2
	var b = 3
	a, b = b, a
	fmt.Println(a, b)
}
