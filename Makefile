
build:
	go build ./pkg

test:
	go test ./pkg

run-example:
	go run examples/server.go