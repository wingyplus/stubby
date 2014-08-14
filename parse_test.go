package stubby

import (
	"os"
	"testing"
)

func eq(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expect %s but was %s", expected, actual)
	}
}

func TestParseYAML(t *testing.T) {
	var (
		f, _      = os.Open("./testdata/helloworld.yaml")
		stubs, err = Parse(f)
	)

	defer f.Close()

	testErrorMustNotNil(t, err)
	testRequest(t, stubs[0].Request)
}

func testErrorMustNotNil(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func testRequest(t *testing.T, req Request) {
	eq(req.Method, "GET")
	eq(req.URL, "/hello-world")
}
