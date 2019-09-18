package filelistening

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const name  = "/list/"

type handleError string

func ( handleError handleError) Error() string  {
	return handleError.OnErrorStr()
}
func ( handleError handleError)OnErrorStr() string  {
	return string(handleError)
}

func FileHandler(writer http.ResponseWriter, request *http.Request) error {
	if isContains:= strings.Contains(request.URL.Path, name);!isContains{
		return  handleError(fmt.Sprintf("the path :%s not contans %s",request.URL.Path,name))
	}
	filetarget:=request.URL.Path[len(name):]
	fmt.Println(" filetarget:",filetarget)
	file, error := os.Open(filetarget)
	if error!=nil {
		return error
	}
	defer file.Close()
	bytes, readError := ioutil.ReadAll(file)
	if readError!=nil{
		return  error
	}
	writer.Write(bytes)
	return nil
}