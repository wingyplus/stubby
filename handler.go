package stubby

import (
	"fmt"
	"github.com/wingyplus/filtr"
	"net/http"
)

type mapHandler map[string]http.Handler

func allowedMethod(method string, h http.Handler) http.Handler {
	switch method {
	case "GET":
		return filtr.GET(h)
	case "POST":
		return filtr.POST(h)
	case "PUT":
		return filtr.PUT(h)
	case "DELETE":
		return filtr.DELETE(h)
	}

	return filtr.GET(h)
}
func newHandler(req Request, res Response) http.Handler {
	var handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		for k, v := range res.Headers {
			w.Header().Set(k, v)
		}
		fmt.Fprintf(w, res.Body)
	}

	return allowedMethod(req.Method, http.HandlerFunc(handlerFunc))
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
