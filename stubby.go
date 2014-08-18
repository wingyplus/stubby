package stubby

import (
	"net/http"
	"os"
)

type Request struct {
	Method string
	URL    string
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

type Stub struct {
	Request
	Response
}

func NewStubbyHandler(path string) (http.Handler, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stubs, err := Parse(f)
	if err != nil {
		return nil, err
	}
	handler := NewMapHandler(stubs)
	return handler, nil
}
