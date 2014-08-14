all: get test

test:
	go test github.com/wingyplus/stubby
	go test github.com/wingyplus/stubby/cmd -f ./testdata/helloworld.yaml -http :8081

get:
	go get -v ./...
