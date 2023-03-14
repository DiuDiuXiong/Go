package handler

import (
	"io"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):] //list/fib.txt -> fib.txt
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
