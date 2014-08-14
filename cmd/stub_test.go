package main

import (
	"flag"
	"testing"
)

func TestParseArgs(t *testing.T) {
	flag.Parse()

	if filename != "./testdata/helloworld.yaml" {
		t.Errorf("expect ./testdata/helloworld.yaml but was %s", filename)
	}

	if addr != ":8081" {
		t.Errorf("expect :8081 but was %s", addr)
	}
}
