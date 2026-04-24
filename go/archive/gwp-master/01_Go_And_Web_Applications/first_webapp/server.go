package main

import (
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World.  读写汉字 - 学中文 😃"))
	// fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
