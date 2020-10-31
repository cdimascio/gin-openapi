.PHONY: version check-v
build:
	go build ./pkg

test:
	go test ./pkg

run-example:
	go run examples/server.go

version: check-v
	echo tagging version: v$(v)
	git tag	v${v}


check-v:
ifndef v
	$(error v is undefined)
endif



