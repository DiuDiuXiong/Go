package infra

import (
	"io"
	"net/http"
)

type Retriever struct{}

/**
1. curl <https address>: To get http content of a web page
2. defer ...: run this command after code finish
*/

func (*Retriever) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // run this line at last

	if bytes, err := io.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
