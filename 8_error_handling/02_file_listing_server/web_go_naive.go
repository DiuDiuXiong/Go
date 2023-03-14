package main

/*
func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):] //list/fib.txt -> fib.txt
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		all, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil) // http://localhost:8888/list/8_error_handling/02_file_listing_server/web.go will give this file
	if err != nil {
		panic(err)
	}

}*/
