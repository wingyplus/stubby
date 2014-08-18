package stubby

import (
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"net/http"
)

func TestNewStubbyHandler(t *testing.T) {
	handler, _ := NewStubbyHandler("./testdata/helloworld.yaml")

	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, _ := http.Get(ts.URL + "/hello-world")

	b, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if string(b) != "hello world" {
		t.Errorf("expect \"hello world\" but was %s", string(b))
	}
}
