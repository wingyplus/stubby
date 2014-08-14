package stubby

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
