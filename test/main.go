// HelloWorld project main.go
package main

//import "time"
import (
	//别名
	wisn_fmt "fmt"
)

//常亮
const PI = 3.14

//变量
var wisn = "Wisn"

//类型申明
type leixin int

//结构体
type gopher struct {
}

//接口
type golang interface{}

//大小写可以控制 常亮 变量 类型 接口 的访问权限
/*函数名首字母小写即为private */
func mainff() {
	wisn_fmt.Println("Hello World!" + wisn)
	wisn_fmt.Println("Hello World!fff")
}
func main() {
	mainff()
	PrintWisn()
	VarTest()
	TestSuan()
	Msize()
	Zhizheng()
	IfTest()
	ForTest()
	SwitchTest()
}
func PrintWisn() {
	wisn_fmt.Print("public func-->PrintWisn")
}
