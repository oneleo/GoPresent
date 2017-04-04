// +build OMIT
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	httpHello()
}

func httpHello() { // OMIT START
	mux := http.NewServeMux()
	mux.HandleFunc("/", fHelloWorld)
	fmt.Printf("Please open: http://127.0.0.1:8080\nTips: Run only on localhost!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func fHelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	_ = r
	return
} // OMIT END
