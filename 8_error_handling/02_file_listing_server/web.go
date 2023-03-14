package main

import (
	handler "Go/8_error_handling/02_file_listing_server/file_listing"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		/**
		If the root url is: http.HandleFunc("/", errWrapper(handler.HandleFileList))
		In function handler, if the url is shorter than len("/list/"), the function itself will have error, use defer recover panic to catch that
		*/
		defer func() {
			if r := recover(); r != nil {

				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s\n", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError // 500 do not know what happened
			}
			http.Error(writer, http.StatusText(code), code) // report to response writer, status not found, with status 404
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handler.HandleFileList))
	// http.HandleFunc("/", errWrapper(handler.HandleFileList)) // try localhost:8888/abc

	err := http.ListenAndServe(":8888", nil) // http://localhost:8888/list/8_error_handling/02_file_listing_server/web.go will give this file
	if err != nil {
		panic(err)
	}

}
