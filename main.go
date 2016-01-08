package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)

	var listen string = os.Getenv("LISTEN")

	if listen == "" {
		listen = ":8080"
	}

	http.ListenAndServe(listen, nil)
	fmt.Printf("Listening on port: %s\n", listen)
}
