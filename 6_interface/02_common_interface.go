package main

import (
	"fmt"
	"io"
	"os"
)

type commonInterface interface {
	fmt.Stringer // contain String() string, act like toString() in java.
	fmt.Scanner  // Scan(), check 2_basic_gramma/5_loops for Scan() example.
	io.ReadWriteCloser
	/**
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	File struct for example
	*/
}

func readandwrite(fl io.ReadWriteCloser) {
	fmt.Println(fl) //assert fl is of type io.ReadWriteCloser
}

func main() {
	f := os.File{}
	readandwrite(&f)
}
