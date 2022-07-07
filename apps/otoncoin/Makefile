test:

	go test -tags=without_main -count=1 -timeout=10s ./...

generate:
	find | grep -e .pb.go | xargs -r rm -v
	go generate ./...

install-tools:
	go install github.com/gogo/protobuf/protoc-gen-gogofaster

.PHONY: test generate
