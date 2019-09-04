package main

import "fmt"

type Point struct {
	x int
	y int
	a,b,c byte
}

func instanceStruct() {
	point := Point{1, 3, 1, 2, 3}
	//{1 3 1 2 3}
	point1 := new(Point)
	//&{0 0 0 0 0}
	//最广泛的方式
	point2 := &Point{1, 3, 1, 2, 3}
	//&{1 3 1 2 3}
	fmt.Println(point)
	fmt.Println(point1)
	fmt.Println(point2)
}

type commend struct {
	name string
	test int
	child *commend
}
func initStruct() {
	//指定字段初始化
	relation:=&commend{
		name:"level1",
		test:1,
		child:&commend{
			name:"level2",
			test:2,
		},
	}
	fmt.Println(relation)
}
func initlist()  {
	//顺序初始化，字段前后有要求
	commend:=&commend{
		"wisn",
		27,
		&commend{},
	}
	fmt.Println(commend)
}
func initNoName() {
	//初始化匿名结构体
	teststruts := &struct {
		name string
		age  int
	}{
		"wisn",
		27,
	}
	fmt.Println(teststruts)

}

func main() {
	printNoNameStruct(&struct {
		id   int
		name string
	}{id:23 , name: "wisn"})
}
func printNoNameStruct(msg *struct{
	id int
	name string
})  {
	fmt.Println("id ",msg.id," name ",msg.name)

}



