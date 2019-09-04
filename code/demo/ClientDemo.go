package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client:=&http.Client{}
	request,error:=	http.NewRequest("GET","https://m.laiyifen.com",strings.NewReader("key=value"))
	if error!=nil{
		fmt.Println(error)
		os.Exit(1)
		return
	}
	request.Header.Add("User-Agent","wisn")
	Response,Reserror:=client.Do(request)
	if Reserror!=nil{
		fmt.Println(Reserror)
		os.Exit(1)
		return
	}

	bytes,error:= ioutil.ReadAll(Response.Body)
	if error!=nil{
		fmt.Println(error)
		os.Exit(1)
		return
	}
	defer Response.Body.Close()
	fmt.Println(string(bytes))
}
