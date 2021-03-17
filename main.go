package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name, err  := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, name)
}
