package main

import "fmt"

type treeNode struct{
	value int
	left ,right *treeNode
}
//传值的，go中所有的都是传值
func (node treeNode)print()  {
	fmt.Println(node.value)
}
func (node *treeNode)setValue(value int)  {
	node.value=value
}
func (node *treeNode) traverse()  {
	if node==nil{
		return ;
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}
//工厂函数
func FactoryCreateNode(v int)*treeNode  {
	return &treeNode{value:v}
}

func main() {
	root:=treeNode{
		value:33}
		root.left=&treeNode{
			value:99,
		}
		root.left.left=new (treeNode)
		root.right=FactoryCreateNode(100)
	fmt.Println(root)
	fmt.Println(root.left)
	nodes:=[]treeNode{
		{value:1},
		{value:2},
		{value:3,left:&root},
	}
	fmt.Println(nodes)
	root.print()
	root.setValue(44)
	root.print()
	root.traverse()
}
