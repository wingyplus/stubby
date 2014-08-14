package stubby

import (
	"os"
	"reflect"
	"testing"
)

func eq(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expect %v but was %v", expected, actual)
	}
}

func TestParseYAML(t *testing.T) {
	var (
		f, _       = os.Open("./testdata/helloworld.yaml")
		stubs, err = Parse(f)
	)

	defer f.Close()

	testErrorMustNotNil(t, err)
	testRequest(t, stubs[0].Request)
	testResponse(t, stubs[0].Response)
}

func testErrorMustNotNil(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func testRequest(t *testing.T, req Request) {
	eq(t, req.Method, "GET")
	eq(t, req.URL, "/hello-world")
}

func testResponse(t *testing.T, res Response) {
	eq(t, res.Status, 200)
	eq(t, res.Headers, map[string]string{"content-type": "application/json"})
	eq(t, res.Body, "hello world")
}
