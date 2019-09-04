package main

import "fmt"
/**
事件方法回调
 */

var eventCallbackMap =make(map[string][]func(interface{}))

func Register(name string,callback func(interface{}))  {
	list:=eventCallbackMap[name]
	list=append(list,callback)
	eventCallbackMap[name]=list
}

func CallEvent(name string,params interface{})  {
	list:=eventCallbackMap[name]
	for _,fun :=range list{
		fun(params)
	}
}

func GlobalCall(params interface{})  {
	fmt.Println("GlobalCall",params)
}

type ViewEvent struct {

}

func (ve *ViewEvent )ViewCall(pointx interface{} )  {
	fmt.Println("ViewCall",pointx)
}
func main() {
	callName:="callview"
	Register(callName,GlobalCall)
	Register(callName,GlobalCall)
	Register(callName,GlobalCall)
	Register(callName,GlobalCall)
	Register(callName,GlobalCall)
	viewevent:=new(ViewEvent)
	Register(callName,viewevent.ViewCall)
	CallEvent(callName,"heheh")
}