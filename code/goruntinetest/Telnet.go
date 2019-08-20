package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ch:=make(chan int)
	TelnetServer("192.168.132.35:3333",ch)
}
func TelnetServer(address string, exit chan int) {
	listener, e := net.Listen("tcp", address)
	if e != nil {
		fmt.Println(e)
		exit <- 1
	}
	fmt.Println("listener:", address)
	defer listener.Close()
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println(e)
			continue
		}
		handleSession(conn, exit)
	}
}
func handleSession(conn net.Conn, exit chan int) {
	fmt.Println("handleSession start")
	reader := bufio.NewReader(conn)
	for{
		s, e := reader.ReadString('\n')
		if e!=nil{
			fmt.Println(e)
			conn.Close()
			exit<-1
		}
		result:=strings.TrimSpace(s)
		if !processCommand(conn,result,exit){
			conn.Write([]byte("wisn:"+result+"\r\n"))
		}
	}
}

func processCommand(conn net.Conn,result string, exit chan int ) bool {
	if strings.HasPrefix(result,"@close"){
		//关闭链接
		conn.Write([]byte("wisn: hehe connect will be close\r\n"))
		conn.Close()
		return true
	}else if strings.HasPrefix(result ,"@shutdown"){
		//关闭服务
		conn.Write([]byte("wisn: hehe server will be close\r\n"))
		exit<-1
		return true
	}
	return false
}
