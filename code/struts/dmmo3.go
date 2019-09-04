package main

import "fmt"

/**
内嵌
 */
type demo1 struct {
	a1 int
	a2 int
}
type demo2 struct {
	demo1
	a3 int
	a4 int
}
//测试内嵌
func testNeiqian() {
	d2 := &demo2{}
	d2.a1 = 1
	d2.a2 = 2
	d2.a3 = 3
	d2.a4 = 4
	fmt.Println(d2)
	//&{{1 2} 3 4}
}

type Fly struct {
}

func (fly *Fly)Flying(name string)  {
	fmt.Println(name,"fly")
}
type Walk struct {
}

func (walk *Walk)walking(name string)  {
	fmt.Println(name,"walk")
}

type People struct {
	Walk
}
type Bird struct {
	Fly
	Walk
}
func ObjectJicheng() {
	people := new(People)
	people.walking("wisn")
	bird := new(Bird)
	bird.walking("aa")
	bird.Flying("bb")
}

func main() {
}


