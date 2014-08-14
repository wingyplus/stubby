package main

import (
	//"github.com/wingyplus/stubby"
	"flag"
)

var (
	filename string
	addr string
)

func init() {
	flag.StringVar(&filename, "f", "", "")
	flag.StringVar(&addr, "http", ":8080", "")
}
