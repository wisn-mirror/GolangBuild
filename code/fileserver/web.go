package main

import (
	"GolangBuild/code/fileserver/filelistening"
	"log"
	"net/http"
	"os"
)
type ExError interface {
	error
	OnErrorStr() string
}

type appHandlerError func(write http.ResponseWriter, request *http.Request) error

func errorWrapper(appHandlerError appHandlerError) func(write http.ResponseWriter, request *http.Request) {
	return func(write http.ResponseWriter, request *http.Request) {
		defer func() {
			if panicError := recover(); panicError != nil {
				http.Error(write, "！！！  server error ", http.StatusInternalServerError)
			}
		}()
		error := appHandlerError(write, request)
		if error != nil {
			log.Printf("error is %v \n", error)
			if exError, ok := error.(ExError); ok {
				http.Error(write, exError.OnErrorStr(), http.StatusInternalServerError)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(error):
				code = http.StatusNotFound
			case os.IsPermission(error):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(write,
				http.StatusText(code),
				http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/", errorWrapper(filelistening.FileHandler))
	serveError := http.ListenAndServe(":9999", nil)
	if serveError != nil {
		panic(serveError)
	}
}
