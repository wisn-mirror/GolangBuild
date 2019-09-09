package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, error := net.Listen("tcp", "localhost:8000")
	if error!=nil{
		log.Print(error)
		return
	}
	go broadCaster()
	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Print(error)
			continue
		}
		go handleconnection(conn)
	}
}

type client chan<- string

var (
	entering=make(chan client)
	leaving=make(chan client)
	messages=make(chan string)
)

func handleconnection(connection net.Conn) {
	ch:=make(chan string)
	go clientWriter(connection,ch)
	who:=connection.RemoteAddr().String()
	ch<-"you are "+ who
	messages<-who+"has arrived"
	entering<-ch
	input:=bufio.NewScanner(connection)
	for input.Scan(){
		messages<-who+":"+input.Text()
	}
	leaving<-ch
	messages<-who+" has left"
	connection.Close()
}
func clientWriter(connection net.Conn,ch<-chan string)  {
	for msg:=range ch {
		fmt.Println(connection,msg)
	}
}

func broadCaster() {
	clients:=make(map[client]bool)
	for{
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case clienter := <-entering:
			clients[clienter] = true
		case clientleaving := <-leaving:
			delete(clients, clientleaving)
			close(clientleaving)
		}
	}

}
