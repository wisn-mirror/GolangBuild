package main

import "fmt"


//构造
func newDemo(name string,age int) *Demo {
	return &Demo{
		name,
		age,
	}
}

func newDemoByName(name string)*Demo  {
	return &Demo{
		name:name,
	}
}

func newDemoTest() {
	fmt.Println(newDemo("wisn", 20))
	fmt.Println(newDemoByName("wisn2"))
}
type Demo struct {
	name string
	age int
}

type DemoChild struct {
	Demo//派生
	height int
}

func newDemoChild(name string)*DemoChild  {
	demochild:=&DemoChild{}
	demochild.name=name
	return demochild
}

type people struct{
	name string
	age int
}
//接收器
func (p *people)setInfo(name string,age int)  {
	p.name=name+"封装一些操作"
	p.age=age
}
func (p *people)getName()string  {
	return p.name
}
func (p *people)setName(name string)  {
	p.name=name
}
func objectDemo() {
	people := &people{}
	people.setInfo("wisn", 23)
	fmt.Println(people)
}

func main() {
}


