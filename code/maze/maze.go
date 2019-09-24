package main

import (
	"fmt"
	"os"
)

func main() {
	 file ,error:= os.Open("maze.in")
	 if error!=nil{
	 	panic(error)
	 }
	 var cow,col int
	 fmt.Fscanf(file,"%d %d" , &cow, &col)
	 fmt.Println(cow,col)
	 maze:=make([][]int ,cow)
	 for index,value:= range maze{
	 	fmt.Println(index ,value,)
		 maze[index]=make([]int ,col)
		 for i:=range maze[index]{
		 	fmt.Fscanf(file,"%d",&maze[index][i])
		 }
		 fmt.Fscanf(file,"%d" , &cow, &col)

	 }
	 for _,value:=range maze{
	 	fmt.Println(value)
	 }




}