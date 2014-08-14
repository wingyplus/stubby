package stubby

type Request struct {
	Method string
	URL    string
}

type Stub struct {
	Request
}
