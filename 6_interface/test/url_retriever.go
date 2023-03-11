package test

import "fmt"

type Retriever struct{}

func (*Retriever) Get(url string) (string, error) {
	return "TEST MESSAGE", nil
}

func (*Retriever) PrintMessage(msg string) {
	fmt.Println(msg)
}
