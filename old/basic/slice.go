package main

import "fmt"

func sliceTest()  {
	array:=[...] int{3,2,5,1,32,6,8,9}
	printAllInfo(array[:])
	temp:=array[:]
	updateSlice(temp)
	printAllInfo(array[2:3])
	printAllInfo(array[2:])
	printAllInfo(array[:8])
	printAllInfo(array[:])
	temp1:=temp[2:6]
	printAllInfo(temp1)
	temp2:=temp1[2:6]
	printAllInfo(temp2)
	appendarray:=[...]int{2,1}
	printAllInfo(appendarray[:])
	//数组添加元素时如果超过cap，系统会重新分配更大的底层数组,策略是2的n次幂
	append1:=append(appendarray[:], 999)
	printAllInfo(append1)
	append2:=append(append1, 999)
	printAllInfo(append2)
	append3:=append(append2, 999)
	printAllInfo(append3)
	append4:=append(append3, 999)
	append5:=append(append4, 999)
	append6:=append(append5, 999)
	append7:=append(append6, 999)
	append8:=append(append7, 999)
	printAllInfo(append8)


}

func updateSlice(array []int )  {
	array[0]=1111
}


func printAllInfo(arr []int)  {
	fmt.Printf("array=%v,len=%d,cap=%d\n",arr,len(arr),cap(arr))
}

func createSlice()  {
	fmt.Println("create")
	var s []int
	printAllInfo(s)
	s1:=[]int{1,5,2,3}
	printAllInfo(s1)
	s2:=make([]int ,100)
	printAllInfo(s2)
	s3:=make([]int ,9,100)
	printAllInfo(s3)
	fmt.Println("copy")
	copy(s3,s1)
	printAllInfo(s3)
	updateSlice(s1)
	printAllInfo(s1)
	printAllInfo(s3)
	s3=append(s1[:3],s3[6:]...)
	printAllInfo(s3)


}
func main() {
	sliceTest()
	createSlice()
}
