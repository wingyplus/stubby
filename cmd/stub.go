package main

import (
	"flag"
	"github.com/wingyplus/stubby"
	"log"
	"os"
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

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	stubs, err := stubby.Parse(f)
	if err != nil {
		panic(err)
	}
	handler := stubby.NewMapHandler(stubs)

	log.Printf("Listen on address %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
