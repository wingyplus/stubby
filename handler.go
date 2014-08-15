package stubby

import (
	"fmt"
	"net/http"
)

type mapHandler map[string]http.Handler

func newHandler(req Request, res Response) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range res.Headers {
			w.Header().Set(k, v)
		}
		fmt.Fprintf(w, res.Body)
	})
}

func NewMapHandler(stubs []Stub) http.Handler {
	var m = make(mapHandler)
	for _, stub := range stubs {
		m[stub.Request.URL] = newHandler(stub.Request, stub.Response)
	}
	return m
}

func (m mapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m[r.URL.Path]; ok {
		h.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}
