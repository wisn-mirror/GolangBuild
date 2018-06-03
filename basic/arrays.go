package basic

import "fmt"

func arraysTest1() {
	var array1 [5]int
	array2 := [2] int{1, 2}
	array3 := [...]int{1, 2, 3}
	var array4 [4][2] bool
	fmt.Println(array1, array2, array3, array4)
}
func arrayTest2() {
	array1 := [4]int{1, 2, 100, 400}
	for i, value := range array1 {
		fmt.Println(i, value)
	}

	for _, value := range array1 {
		fmt.Println(value)
	}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
}

func printArray(arr *[7]int) {
	arr[0]=100
	for i, value := range arr {
		fmt.Println(i, value)
	}
}
func printArray2(arr []int)  {
	fmt.Println("print-------------start")
	for i,value:=range arr{
		fmt.Println(i,value)
	}
	fmt.Println("print-------------end")
}
func main() {
	//arraysTest1()
	//arrayTest2()
	array1:=[7]int{1,3,4,2,7,6,9}
	printArray(&array1)
	printArray2(array1[:])

}
