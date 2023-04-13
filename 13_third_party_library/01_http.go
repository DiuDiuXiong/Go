package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// for http server check 8_error_handling

func request(url string) {
	req, err := http.NewRequest("GET", url, nil)
	// defer req.Body.Close()
	if err != nil {
		panic(err)
	}

	// can change to different user agent, by doing:
	// 1. open desired website on chrome
	// 2. Change to other device version
	// 2. Inspect/Network/Request Headers/User Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request,
		) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

// Can also use pprof to verify which function consume more time

/*
func main() {
	request("https://imooc.com")
}
*/
