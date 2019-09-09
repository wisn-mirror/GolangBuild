package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	connection ,error:=net.Dial("tcp","localhost:8000")
	if error!=nil{
		log.Print(error)
	}
	done:=make(chan struct{})
	go func() {
		io.Copy(os.Stdout,connection)
		log.Println("done")
		done<- struct{}{}
	}()
	mustcopy(connection,os.Stdin)
	connection.Close()
	<-done
}

func mustcopy(conn net.Conn, src io.Reader) {
	if _,error:=io.Copy(conn,src);error!=nil{
		log.Fatal(error)
	}
}
