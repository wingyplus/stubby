package main

import (
	"flag"
	"github.com/wingyplus/stubby"
	"log"
	"net/http"
)

var (
	filename string
	addr     string
)

func init() {
	flag.StringVar(&filename, "f", "", "")
	flag.StringVar(&addr, "http", "0.0.0.0:8080", "")
}

func main() {
	flag.Parse()

	handler, err := stubby.NewStubbyHandler(filename)
	if err != nil {
		panic(err)
	}

	log.Printf("Listen on address %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
