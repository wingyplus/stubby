package stubby

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	var (
		req = Request{
			Method: "GET",
			URL:    "/hello-world",
		}
		res = Response{
			Status: 200,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body: "Hello World",
		}
	)

	var handler http.Handler = NewHandler(req, res)

	var (
		request, _ = http.NewRequest("GET", "/hello-world", nil)
		recorder   = httptest.NewRecorder()
	)

	handler.ServeHTTP(recorder, request)

	if recorder.Body.String() != "Hello World" {
		t.Errorf("expect Hello World but was %s", recorder.Body.String())
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("expect status 200 but was %d", recorder.Code)
	}

	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expect application/json but was %s", recorder.Header().Get("Content-Type"))
	}
}

func TestNewMapHandler(t *testing.T) {
	var stubs []Stub = []Stub{
		Stub{
			Request{
				Method: "GET",
				URL:    "/hello-world",
			},
			Response{
				Status: 200,
				Headers: map[string]string{
					"content-type": "application/json",
				},
				Body: "Hello World",
			},
		},
		Stub{
			Request{
				Method: "GET",
				URL:    "/hello-xml",
			},
			Response{
				Status: 200,
				Headers: map[string]string{
					"content-type": "text/xml",
				},
				Body: "<a>Hello</a>",
			},
		},
	}

	var handler http.Handler = NewMapHandler(stubs)

	var (
		request, _ = http.NewRequest("GET", "http://localhost/hello-world", nil)
		recorder   = httptest.NewRecorder()
	)

	handler.ServeHTTP(recorder, request)

	if recorder.Body.String() != "Hello World" {
		t.Errorf("expect Hello World but was %s", recorder.Body.String())
	}

	request, _ = http.NewRequest("GET", "http://localhost/hello-xml", nil)
	recorder = httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Body.String() != "<a>Hello</a>" {
		t.Errorf("expect <a>Hello</a> but was %s", recorder.Body.String())
	}
}
